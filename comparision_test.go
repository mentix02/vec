package vec

import "testing"

func TestVec_EqualsVec(t *testing.T) {
	i := []int64{1, 2, 3, 4, 5}
	v := NewFromInt(i)
	ov := NewFromInt(i)
	if !v.EqualsVec(ov) {
		t.Errorf("%v != %v", v.Range(), ov.Range())
	}
	_, _ = ov.Pop()
	if v.EqualsVec(ov) {
		t.Errorf("%v == %v", v.Range(), ov.Range())
	}
	_, _ = v.Pop()
	v.Set(2, 10)
	if v.EqualsVec(ov) {
		t.Errorf("%v == %v", v.Range(), ov.Range())
	}

	f := []float64{1.31, 3.14, 7.31, 9.1,3}
	v = NewFromFloat(f)
	fov := NewFromFloat(f)
	if !v.EqualsVec(fov) {
		t.Errorf("%v != %v", v.Range(), fov.Range())
	}
	v.Set(2, 14.3)
	if v.EqualsVec(fov) {
		t.Errorf("%v == %v", v.Range(), fov.Range())
	}

	if NewFromInt(i).EqualsVec(NewFromFloat(f)) {
		t.Errorf("mismatched numeric type vectors equal")
	}

}

func TestVec_EqualsVec3(t *testing.T) {
	s := []interface{}{"a", "b", "c"}
	v := NewFromSlice(s)
	ov := NewFromSlice(s)
	if !v.EqualsVec(ov) {
		t.Errorf("%v != %v", v.Range(), ov.Range())
	}
	v.Set(2, "d")
	if v.EqualsVec(ov) {
		t.Errorf("%v == %v", v.Range(), ov.Range())
	}
}

func TestVec_Swap(t *testing.T) {
	v := NewFromInt([]int64{1, 2, 3, 4})
	err := v.Swap(0, 3)

	if err != nil {
		t.Errorf(err.Error())
	}

	if v.AtInt(0) != int64(4) && v.AtInt(3) != int64(1) {
		t.Error("swap failed")
	}

	err = v.Swap(4, 10)
	if err == nil {
		t.Error("out of bounds swap committed")
	}
}
