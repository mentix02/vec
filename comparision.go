package vec

import "errors"

func (v *Vec) Swap(i1, i2 uint64) error {
	if i1 > v.Len() || i2 > v.Len() {
		return errors.New("list index out of range")
	} else {
		v.data[i1], v.data[i2] = v.data[i2], v.data[i1]
		return nil
	}
}

//EqualsVec checks for equality between two vectors.
func (v *Vec) EqualsVec(ov *Vec) bool {
	if v.Len() != ov.Len() {
		return false
	} else if !(v.integer == ov.integer && v.floating == ov.floating) {
		return false
	} else {
		if v.integer {
			for index, item := range v.Range() {
				if ov.AtInt(uint64(index)) != item {
					return false
				}
			}
		} else if v.floating {
			for index, item := range v.Range() {
				if ov.AtFloat(uint64(index)) != item {
					return false
				}
			}
		} else {
			for index, item := range v.Range() {
				if ov.At(uint64(index)) != item {
					return false
				}
			}
		}
		return true
	}
}
