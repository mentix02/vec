package vec

import (
	"errors"
	"fmt"
)

// Range returns data sliced until the length of the vector. This is
// done so that the unexported field Vec.data is not used manually.
func (v *Vec) Range() []interface{} {
	return v.data[:v.Len()]
}

// At returns the value stored at the index provided.
func (v *Vec) At(index uint64) interface{} {
	return v.data[index]
}

// SafeAt returns the value stored at the index provided
// with some elementary bounds and thread safety checks.
func (v *Vec) SafeAt(index uint64) (interface{}, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if index >= v.Len() {
		return nil, errors.New("index out of bounds")
	} else {
		return v.data[index], nil
	}
}

// Set sets the value stored at the index provided.
func (v *Vec) Set(index uint64, item interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.data[index] = item
}

// SetInt sets the value stored at the index provided as an int64.
func (v *Vec) SetInt(index uint64, item interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.data[index] = convertToInt64(item)
}

// SetFloat sets the value stored at the index provided as a float64.
func (v *Vec) SetFloat(index uint64, item interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.data[index] = convertToFloat64(item)
}

// AtInt returns the integer value stored at the index provided.
func (v *Vec) AtInt(index uint64) int64 {
	return convertToInt64(v.data[index])
}

// AtFloat returns the floating value stored at the index provided.
func (v *Vec) AtFloat(index uint64) float64 {
	return convertToFloat64(v.data[index])
}

// First returns a pointer pointing to the first member of the vector.
func (v *Vec) First() *interface{} {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.Len() > 0 {
		return &v.Range()[0]
	}
	return nil
}

// Last returns a pointer pointing to the last member of the vector.
func (v *Vec) Last() *interface{} {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.Len() > 0 {
		return &v.Range()[v.Len()-1]
	}
	return nil
}

// Index is similar to Contains as in it performs as linear search
// but instead of returning a boolean, it returns the index if it's
// found else returns an error.
func (v *Vec) Index(toFind interface{}) (uint64, error) {
	if v.integer {
		for index, item := range v.Range() {
			if item == convertToInt64(toFind) {
				return uint64(index), nil
			}
		}
	} else if v.floating {
		for index, item := range v.Range() {
			if item == convertToFloat64(toFind) {
				return uint64(index), nil
			}
		}
	} else {
		for index, item := range v.Range() {
			if item == toFind {
				return uint64(index), nil
			}
		}
	}
	return 0, fmt.Errorf("'%v' not found in vector", toFind)
}

// Copy clones the data of a vector with data such as the capacity,
// range, length, etc (but not the mutex) into a new vector.
func (v *Vec) Copy() *Vec {
	n := New(v.capacity)
	for i := range n.data {
		n.data[i] = v.Range()[i]
	}
	n.length = v.Len()
	return n
}

// ToInt64Slice converts a vector into a slice of int64s.
func (v *Vec) ToInt64Slice() ([]int64, error) {
	if !v.integer {
		return nil, errors.New("numeric vector expected")
	}
	s := make([]int64, v.Len(), v.Len())
	for index, item := range v.Range() {
		s[index] = item.(int64)
	}
	return s, nil
}

// ToFloat64Slice converts a vector into a slice of int64s.
func (v *Vec) ToFloat64Slice() ([]float64, error) {
	if !v.floating {
		return nil, errors.New("numeric vector expected")
	}
	s := make([]float64, v.Len(), v.Len())
	for index, item := range v.Range() {
		s[index] = item.(float64)
	}
	return s, nil
}
