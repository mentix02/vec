package vec

import "testing"

func TestVec_EqualsVec(t *testing.T) {
	v := NewFromInt([]int64{1, 2, 3, 4, 5})
	ov := NewFromInt([]int64{1, 2, 3, 4, 5})
	if !v.EqualsVec(ov) {
		t.Errorf("%v != %v", v.Range(), ov.Range())
	}
	_, _ = ov.Pop()
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
