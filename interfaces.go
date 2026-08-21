// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer

//*** interfaces to the RingBuffer
// * As a [bounded] queue
// * As a native ring buffer

// Queue represents the standard queue operations
type Queue[T comparable] interface {
	Enqueue(item T) error
	Dequeue() (T, error)
	// Cap returns the outer capacity of the ring buffer.
	Cap() BufSizeT
	// Len returns the quantity of items in the [ring buffer] queue
	Len() BufSizeT
	// Reset empties+reinitializes the buffer
	Reset()

	IsEmpty() bool
	IsFull() bool
}

///////////////////////////////////////////

// RingBuffer - set of standard ring buffer operations (includes Queue!)
type RingBuffer[T comparable] interface {
	Close()

	// Queue[T]
	Enqueue(item T) error         // equiv to [Put]
	Dequeue() (item T, err error) // equiv to [Get]
	Cap() BufSizeT                // outer capacity of the ring buffer.
	CapReal() BufSizeT            // real (inner) capacity of the ring buffer.
	Len() BufSizeT                // count of items in the queue
	Reset()
	IsEmpty() bool
	IsFull() bool
}

// SimpleRingBuffer is the simplest interface to a (generic) Ring Buffer
type SimpleRingBuffer[T comparable] interface {
	////// aliases of functions in RingBuffer
	Put(item T) error // Put item into RB
	Get() (T, error)  // Retrieve item from RB
	Count() int       // Count of items in the RB queue
	Cap() int
	IsEmpty() bool
	IsFull() bool
}
