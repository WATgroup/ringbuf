// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer_test

import (
	"fmt"
	rbp "oss.w-a-t.group/ringbuf"
	"testing"
)

type myPage = [64]byte

func TestPage(t *testing.T) {

	var rb rbp.SimpleRingBuf[myPage]
	rb.Init(3)

	fmt.Println("cap[empty]=>", rb.Cap())
	fmt.Println("count[empty]=>", rb.Count())

	for j := range rb.Cap() {
		pp, _ := rb.PeekWrite()
		// fmt.Println("pp=>", pp)
		fillPage(pp, byte(32*j)+'a')
		_ = rb.Push()
		// fmt.Println("count=>", rb.Count())
	}

	fmt.Println("cap[full]=>", rb.Cap())
	fmt.Println("count[full]=>", rb.Count())

	// fmt.Println("-------------------------------------------------------------------------")
	// fmt.Println(rb)
	// fmt.Println("-------------------------------------------------------------------------")

	for i := range rb.Count() {
		v, _ := rb.Get()
		fmt.Println(i, "\t", v)
	}

	fmt.Println("cap[full]=>", rb.Cap())
	fmt.Println("count[full]=>", rb.Count())
}

func fillPage(x *myPage, start byte) {
	// fmt.Printf("page=>%p: %q\n", x, x)
	for i := range 64 {
		x[i] = (start + byte(i))
	}
}
