package vec

import (
	"errors"
	"math/rand"
)

func quicksort(data []interface{}, integer bool) []interface{} {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1

	pivot := rand.Int() % len(data)

	data[pivot], data[right] = data[right], data[pivot]

	if integer {
		for i := range data {
			if data[i].(int64) < data[right].(int64) {
				data[left], data[i] = data[i], data[left]
				left++
			}
		}
	} else {
		for i := range data {
			if data[i].(float64) < data[right].(float64) {
				data[left], data[i] = data[i], data[left]
				left++
			}
		}
	}

	data[left], data[right] = data[right], data[left]

	quicksort(data[:left], integer)
	quicksort(data[left+1:], integer)

	return data
}

// Sort sorts the data of a vector if it's of integer type.
func (v *Vec) Sort() bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	if !(v.integer || v.floating) {
		return false
	} else {
		copy(v.data[:v.Len()-1], quicksort(v.Range(), v.integer))
		return true
	}
}

// IsSorted checks whether a given numeric vector is sorted.
func (v *Vec) IsSorted() (bool, error) {
	var index uint64
	if v.integer {
		for index = 0; index < v.Len()-1; index++ {
			if v.AtInt(index) > v.AtInt(index+1) {
				return false, nil
			}
		}
		return true, nil
	} else if v.floating {
		for index = 0; index < v.Len()-1; index++ {
			if v.AtFloat(index) > v.AtFloat(index+1) {
				return false, nil
			}
		}
		return true, nil
	}
	return false, errors.New("expecting a numeric vector")
}
