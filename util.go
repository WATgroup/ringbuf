// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer

const minRBsize = 4

func roundUp(n BufSizeT) uint32 {
	n = max(n, minRBsize)
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}

func absDiff(l, r BufSizeT) BufSizeT {
	// if l >= r {
	// 	return (l-r)
	// } else {
	// 	return -(l-r)
	// }
	if l < r {
		return -(l - r)
	}
	return (l - r)
}

// func main() {
// 	fmt.Println(roundUp(15))
// 	fmt.Println(roundUp(23))
// 	fmt.Println(roundUp(1))
// 	fmt.Println(roundUp(126))
// 	fmt.Println(roundUp(254))
// 	fmt.Println(roundUp(257))
// 	fmt.Println(roundUp(4094))
// 	fmt.Println(roundUp(8191))
// 	fmt.Println(roundUp(32767))
// }
