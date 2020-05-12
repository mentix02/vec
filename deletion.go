package vec

import "errors"

// Pop removes the last element from a vector and returns it.
func (v *Vec) Pop() (interface{}, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.Len() < 1 {
		return nil, errors.New("pop from empty vector")
	} else {
		item := v.data[v.Len()-1]
		v.data[v.Len()-1] = nil
		v.length--
		return item, nil
	}
}

// Clear deleted all items from a Vector and
// allocates Vector.Data to new 0 byte long memory chunk.
func (v *Vec) Clear() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.length, v.capacity = 0, 0
	v.data = make([]interface{}, 0, 0)
}

// Remove removes the first occurrence of item provided.
func (v *Vec) Remove(item interface{}) bool {
	v.mu.Lock()
	index, err := v.Index(item)
	if err == nil {
		copy(v.data[index:], v.data[index+1:])
		v.data[v.Len()-1] = nil
		v.length--
		v.data = v.data[:v.capacity-1]
		v.mu.Unlock()
		return true
	}
	v.mu.Unlock()
	return false
}
