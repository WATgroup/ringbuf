// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package ringbuffer

import (
	"github.com/WATgroup/errors"
)

var (
	errRBnoinit = errors.New("RingBuffer not initialized")
	errRBinited = errors.New("RingBuffer already initialized!")
	errRBempty  = errors.New("RingBuffer is empty")
	errRBfull   = errors.New("RingBuffer is full")
)
