package vec

import "testing"

func TestVec_At(t *testing.T) {
	s := []int64{3, 1, 4, 6}
	v := NewFromInt(s)
	for index, item := range s {
		if v.At(uint64(index)) != item {
			t.Errorf("v.At(%d) (%d) != %d", index, v.At(uint64(index)), item)
		}
	}
}

func TestVec_AtFloat(t *testing.T) {
	s := []float64{3.14, 7.53, 41.31, 41.3}
	v := NewFromFloat(s)
	for index, item := range s {
		if v.AtFloat(uint64(index)) != item {
			t.Errorf("v.AtFloat(%d) (%f) != %f", index, v.AtFloat(uint64(index)), item)
		}
	}
}

func TestVec_Copy(t *testing.T) {
	s := []int64{3, 1, 4, 6}
	v := NewFromInt(s)
	ov := v.Copy()
	if ov.Len() != v.Len() {
		t.Errorf("ov.Len() [%d] != v.Len() [%d]", ov.Len(), v.Len())
	}
}
