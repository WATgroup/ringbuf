// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer

// SimpleRingBuf implements a [fixed size] *single-threaded* circular buffer:
// new writes will be rejected when it becomes full.
type SimpleRingBuf[T comparable] struct {
	data  []T     // data elements
	flags RBflags // rb flags, mostly for completitude
	//...also, for alignment
	bufCap uint32
	head   uint32 // writeCursor
	tail   uint32 // readCursor
	cnt    uint32
	_      [cacheLinePad - 16]byte //nolint:revive
	// *** rbItem is already 'cacheLinePad' padded...
	helper RBItem[T]
}

var _ = (SimpleRingBuffer[rune])((*SimpleRingBuf[rune])(nil))

func NewRingBuf[T comparable](size BufSizeT) *SimpleRingBuf[T] {
	sz := roundUp(size)
	ret := &SimpleRingBuf[T]{data: make([]T, sz),
		bufCap: sz, // rounded up to power-of-2
	}
	// head,tail already initialized to 0
	return ret
}

func (rb *SimpleRingBuf[T]) Init(size BufSizeT) error {
	if nil != rb.data {
		return errRBinited
	}
	sz := roundUp(size)
	rb.data = make([]T, sz)
	rb.bufCap = sz
	return nil
}

func (rb *SimpleRingBuf[T]) Reset() error {
	if nil == rb.data {
		return errRBnoinit
	}
	rb.head = 0
	rb.tail = 0
	rb.cnt = 0
	rb.flags = 0
	return nil
}

// SetHelper registers the item BlT helper
func (rb *SimpleRingBuf[T]) SetHelper(hh RBItem[T]) {
	if nil == rb.data {
		return // errRBnoinit
	}
	rb.helper = hh
}

func (rb *SimpleRingBuf[T]) Put(item T) error {
	if nil == rb.data {
		return errRBnoinit
	}
	rb.data[rb.head] = item
	rb.head = (rb.head + 1) & (rb.bufCap - 1)
	rb.cnt++
	return nil
}

func (rb *SimpleRingBuf[T]) Get() (item T, err error) {
	if nil == rb.data {
		err = errRBnoinit
		return
	}
	item = rb.data[rb.tail]
	rb.tail = (rb.tail + 1) & (rb.bufCap - 1)
	rb.cnt--
	return
}

func (rb *SimpleRingBuf[T]) Count() int {
	if nil == rb.data {
		return -1
	}
	// return int(absDiff(rb.head, rb.tail))
	return int(rb.cnt)
}

func (rb *SimpleRingBuf[T]) Cap() int {
	if nil == rb.data {
		return -1
	}
	return cap(rb.data)
}

func (rb *SimpleRingBuf[T]) IsEmpty() bool {
	if nil == rb.data {
		return false
	}
	return rb.isEmpty()
}

func (rb *SimpleRingBuf[T]) IsFull() bool {
	if nil == rb.data {
		return false
	}
	return rb.isFull()
}

///////////////////////////////////////////////////////////////////////////

func (rb *SimpleRingBuf[T]) PeekWrite() (*T, error) {
	if nil == rb.data {
		return nil, errRBnoinit
	}
	// No checks for emptyness -- we might be preparing the first write
	if rb.isFull() {
		return nil, errRBfull
	}
	return &rb.data[rb.head], nil
}

func (rb *SimpleRingBuf[T]) PeekRead() (*T, error) {
	if nil == rb.data {
		return nil, errRBnoinit
	}
	// No checks for fullness. It's probably Ok
	if rb.isEmpty() {
		return nil, errRBempty
	}
	return &rb.data[rb.tail], nil
}

func (rb *SimpleRingBuf[T]) Push() error {
	if nil == rb.data {
		return errRBnoinit
	}
	if rb.isFull() {
		return errRBfull
	}
	rb.head = (rb.head + 1) & (rb.bufCap - 1)
	rb.cnt++
	return nil
}

func (rb *SimpleRingBuf[T]) Pop() error {
	if nil == rb.data {
		return errRBnoinit
	}
	if rb.isEmpty() {
		return errRBempty
	}
	rb.tail = (rb.tail + 1) & (rb.bufCap - 1)
	rb.cnt--
	return nil
}

func (rb *SimpleRingBuf[T]) isEmpty() bool {
	return 0 == rb.cnt && (rb.tail == rb.head)
}
func (rb *SimpleRingBuf[T]) isFull() bool {
	// return ((rb.tail + 1) & (rb.bufCap - 1)) == rb.head
	return rb.cnt == rb.bufCap
}
