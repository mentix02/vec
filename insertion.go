package vec

import "sync/atomic"

// Append appends a new item to the vector. Append is also
// thread safe. A few notes about Append - when appending
// to a numeric vector, either floating point or integer type,
// Append will assert types and convert if integer type vector
// to int64s for convenience and floats to float64.
//
// It doesn't impact the type for non numeric vectors.
func (v *Vec) Append(item interface{}) {
	v.mu.Lock()
	if !v.full() {
		setToNumericType(v, item)
		atomic.AddUint64(&v.length, 1)
		v.mu.Unlock()
	} else {
		largerCap := largerCapacity(v.capacity)
		data := make([]interface{}, largerCap, largerCap)
		copy(data, v.data)
		v.data = data
		v.capacity = largerCap
		v.mu.Unlock()
		v.Append(item)
		return
	}
}

// Insert inserts an item into a provided index.
func (v *Vec) Insert(index uint64, item interface{}) {
	v.mu.Lock()
	if (index == v.Len() && index <= v.capacity) || index > v.Len() {
		v.mu.Unlock()
		v.Append(item)
		return
	}
	if v.Len()+1 < v.capacity {
		atomic.AddUint64(&v.length, 1)
		for i := v.Len(); i > index; i-- {
			v.data[i] = v.At(i - 1)
		}
		setToNumericType(v, item)
	} else {
		largerCap := largerCapacity(v.capacity)
		largerData := make([]interface{}, largerCap, largerCap)

		copy(largerData, v.data[0:index])
		largerData[index] = item
		for i := index + 1; i <= v.Len(); i++ {
			largerData[i] = v.At(i - 1)
		}

		atomic.AddUint64(&v.length, 1)
		v.data = largerData
		v.capacity = largerCap
		v.mu.Unlock()
	}
}

// InsertFromSlice appends items from a slice to vector.
func (v *Vec) ExtendFromSlice(s []interface{}) {
	for _, item := range s {
		v.Append(item)
	}
}

// InsertFromVec, similar to InsertFromSlice, appends items
// from another vector to the same vector.
func (v *Vec) ExtendFromVec(o *Vec) {
	for _, item := range o.Range() {
		v.Append(item)
	}
}
