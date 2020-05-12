package vec

import "testing"

func TestVec_Pop(t *testing.T) {
	s := []int64{3, 1, 4, 6}
	v := NewFromInt(s)
	for index := len(s) - 1; index >= 0; index-- {
		if val, err := v.Pop(); val.(int64) != s[index] {
			if err != nil {
				t.Errorf(err.Error())
			} else {
				t.Errorf("%d (%T) != %d (%T)", val, val, s[index], s[index])
			}
		}
	}
	_, err := v.Pop()
	if err == nil {
		t.Errorf("items not removed from vector: %v", v.Range())
	}
}

func TestVec_Clear(t *testing.T) {
	v := NewFromInt([]int64{3, 1, 4, 5})
	v.Clear()
	if v.Len() != 0 {
		t.Errorf("vector not cleared completely: %v (len: %d)", v.Range(), v.Len())
	}
}

func TestVec_Remove(t *testing.T) {
	v := NewFromInt([]int64{9, 3, 1, 8, 3})
	res := NewFromInt([]int64{9, 1, 8, 3})
	v.Remove(3)
	v.Remove(-3)
	if !v.EqualsVec(res) {
		t.Errorf("vectors differ: %v != %v", v.Range(), res.Range())
	}
}
