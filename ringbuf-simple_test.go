// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer_test

import (
	"fmt"
	rbp "oss.w-a-t.group/ringbuf"
	"testing"
)

func TestBasics(t *testing.T) {

	rb := rbp.NewRingBuf[int](7) // gets converted to 8
	fmt.Println("cap[empty]=>", rb.Cap())
	fmt.Println("count[empty]=>", rb.Count())

	for i := range rb.Cap() {
		_ = rb.Put(i + 1)
	}
	fmt.Println("count[some]=>", rb.Count())

	for !rb.IsEmpty() {
		x, e := rb.Get()
		if nil != e {
			t.Fail() // e.Error())
		}
		fmt.Println(x)
	}
}

func TestByVal(t *testing.T) {

	var rb rbp.SimpleRingBuf[rune]
	rb.Init(11) // gets converted to 16
	fmt.Println("cap[empty]=>", rb.Cap())
	fmt.Println("count[empty]=>", rb.Count())

	for i := range 13 {
		_ = rb.Put(rune('0' + i + 1))
	}
	fmt.Println("count[some]=>", rb.Count())

	for !rb.IsEmpty() {
		x, e := rb.Get()
		if nil != e {
			t.Fail() // e.Error())
		}
		fmt.Printf("%d - %c\n", x, x)
	}
}
