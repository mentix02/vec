package vec

import "testing"

func TestVec_Sort(t *testing.T) {
	v := NewFromFloat([]float64{4.31, 1.21, 6.92, 4.22, 2.33})
	v.Sort()
	if sorted, err := v.IsSorted(); !sorted || err != nil {
		t.Errorf("vector is still unsorted: %v", v.Range())
	}
	v = NewFromInt([]int64{4, 1, 6, 4, 2})
	v.Sort()
	if sorted, err := v.IsSorted(); !sorted || err != nil {
		t.Errorf("vector is still unsorted: %v", v.Range())
	}
}

func TestVec_IsSorted(t *testing.T) {
	v := NewFromInt([]int64{1, 2, 3, 4, 5, 6})
	if ok, err := v.IsSorted(); !ok || err != nil {
		t.Errorf("%v is not sorted", v.Range())
		if err != nil {
			t.Error(err.Error())
		}
	}
	v = NewFromFloat([]float64{1.21, 2.33, 4.22, 4.31, 6.92})
	if ok, err := v.IsSorted(); !ok || err != nil {
		t.Errorf("%v is not sorted", v.Range())
		if err != nil {
			t.Error(err.Error())
		}
	}
}
