// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtype

import (
	"math"
	"sync/atomic"
	"unsafe"
)

type Float32 struct {
	value uint32
}

// NewFloat32 returns a concurrent-safe object for float32 type,
// with given initial value <value>.
func NewFloat32(value...float32) *Float32 {
    if len(value) > 0 {
        return &Float32{
	        value : math.Float32bits(value[0]),
		}
    }
    return &Float32{}
}

// Clone clones and returns a new concurrent-safe object for float32 type.
func (v *Float32) Clone() *Float32 {
    return NewFloat32(v.Val())
}

// Set atomically stores <value> into t.value and returns the previous value of t.value.
func (v *Float32) Set(value float32) (old float32) {
    return math.Float32frombits(atomic.SwapUint32(&v.value, math.Float32bits(value)))
}

// Val atomically loads t.value.
func (v *Float32) Val() float32 {
    return math.Float32frombits(atomic.LoadUint32(&v.value))
}

// Add atomically adds <delta> to t.value and returns the new value.
func (v *Float32) Add(delta float32) (new float32) {
	for {
		old := math.Float32frombits(v.value)
		new  = old + delta
		if atomic.CompareAndSwapUint32(
			(*uint32)(unsafe.Pointer(&v.value)),
			math.Float32bits(old),
			math.Float32bits(new),
		) {
			break
		}
	}
	return
}

// Cas executes the compare-and-swap operation for value.
func (v *Float32) Cas(old, new float32) bool {
	return atomic.CompareAndSwapUint32(&v.value, uint32(old), uint32(new))
}
