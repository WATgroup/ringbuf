// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

// Package ringbuffer implements two variants of a generic ring buffer
// - Regular / single-threaded
// - Concurrent-safe (multiple producers & consumers)
package ringbuffer

type BufSizeT = uint32

//

type RBflags uint64

const (
	RB_Overlapped RBflags = 1 << (31 - iota)
	RB_Full
)
