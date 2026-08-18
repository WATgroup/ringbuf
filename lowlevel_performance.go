// SPDX-FileCopyrightText: © 2025,2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//go:build performance

package ringbuffer

import (
	"golang.org/x/sys/cpu"
	"unsafe"
)

// cacheLinePad equals the CPU Cache Line Padding Size, ref. current CPU
const cacheLinePad = unsafe.Sizeof(cpu.CacheLinePad{})

func init() {
	metricsEnabled = true
}
