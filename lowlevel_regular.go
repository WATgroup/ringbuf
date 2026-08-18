// SPDX-FileCopyrightText: © 2025,2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//go:build !performance

package ringbuffer

import ()

// cacheLinePad equals the CPU Cache Line Padding Size, ref. current CPU
// when we don't care/favor lower memory usage, hardwire it to 16, which is "big enough"
const cacheLinePad = 16
