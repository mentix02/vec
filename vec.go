// Package vec provides a strong dynamic slice interface for storing
// and updating data efficiently with built in thread safety.
package vec

import (
	"reflect"
	"sync"
)

type Vec struct {
	integer  bool
	floating bool

	length   uint64
	capacity uint64

	mu sync.Mutex

	data []interface{}
}

func (v *Vec) full() bool {
	return v.length == v.capacity
}

// Len returns the number of items in the vector.
func (v *Vec) Len() uint64 {
	return v.length
}

// New returns a generic vector of some capacity.
func New(capacity uint64) *Vec {
	v := &Vec{capacity: capacity, length: 0}
	v.data = make([]interface{}, capacity, capacity)
	return v
}

// NewInt returns an empty vector with specific
// int64 optimizations and feature functions.
func NewInt(capacity uint64) *Vec {
	v := New(capacity)
	v.integer = true
	return v
}

// NewFloat returns an empty vector with specific
// float64 optimizations and feature functions.
func NewFloat(capacity uint64) *Vec {
	v := New(capacity)
	v.floating = true
	return v
}

// Empty returns an empty vector. It can be provided an
// optional reflect.Kind argument that can be used to
// construct an int64 or float64 vector otherwise it defaults
// to an empty interface{} vector.
func Empty(t ...reflect.Kind) *Vec {
	if len(t) > 0 {
		if t[0] == reflect.Int64 {
			return NewFromInt([]int64{})
		} else if t[0] == reflect.Float64 {
			return NewFromFloat([]float64{})
		}
	}
	return New(0)
}

// NewFromSlice constructs a generic Vector from a given slice of interface{}s.
func NewFromSlice(slice []interface{}) *Vec {
	v := New(uint64(len(slice)))
	for _, item := range slice {
		v.Append(item)
	}
	return v
}

// NewFromInt constructs a numeric Vector from a given slice of int64s.
func NewFromInt(slice []int64) *Vec {
	v := New(uint64(len(slice)))
	v.integer = true
	for _, item := range slice {
		v.Append(convertToInt64(item))
	}
	return v
}

// NewFromFloat constructs a numeric Vector from a given slice of float64s.
func NewFromFloat(slice []float64) *Vec {
	v := New(uint64(len(slice)))
	v.floating = true
	for _, item := range slice {
		v.Append(item)
	}
	return v
}

// Resize either deletes elements from the vector if the capacity provided
// is less than the current length else appends more capacity to the vector.
func (v *Vec) Resize(capacity uint64) {
	if capacity < v.Len() {
		v.data = v.data[:capacity]
	} else {
		newData := make([]interface{}, capacity)
		copy(newData, v.data)
		v.data = newData
	}
}
