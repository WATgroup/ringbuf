// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer

// import "unsafe"

// An RBItem supports more efficient clone operations
type RBItem[T comparable] interface {
	PreAlloc(numItems BufSizeT) (newBlock T)
	CloneIn(srcBlock T, targetBlock *T)
	CloneOut(srcBlock *T) (targetBlock T)
}

// the rbItem type is a RingBuffer *item*
type rbItem[T comparable] struct {
	value T      // DATA
	flags uint64 // bits: 0=writable, 1=readable, 2=write ok, 3=read ok
	_     [cacheLinePad - 8 - /*unsafe.Sizeof(*(*T)(nil))*/ 8]byte
}

type rbItemData[T comparable] struct {
	value T // DATA
}

type rbItemMeta[T comparable] struct {
	item  *T
	flags uint64 // bits: 0=writable, 1=readable, 2=write ok, 3=read ok
	_     [cacheLinePad - 8 - 8]byte
}
