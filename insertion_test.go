package vec

import (
	"reflect"
	"testing"
	"time"
)

func TestVec_Append(t *testing.T) {
	s := []int64{2, 1, 4, 3}
	v := NewFromInt(s)
	v.Append(int64(8))
	s = append(s, 8)
	for index, item := range v.Range() {
		if item != s[index] {
			t.Errorf("%d (%T) != %d (%T)", item, item, s[index], s[index])
		}
	}
}

func TestVec_Append2(t *testing.T) {
	s := []int64{2, 1, 4, 3}
	v := Empty(reflect.Int64)
	for _, item := range s {
		go v.Append(item)
	}
	for v.Len() != 4 {
		time.Sleep(5 * time.Millisecond)
	}
	v.Sort()
	for index, item := range []int64{1, 2, 3, 4} {
		if v.AtInt(uint64(index)) != item {
			t.Errorf("vector not sorted or appended right: %v", v.Range())
			break
		}
	}
}

func TestVec_ExtendFromSlice(t *testing.T) {
	s := []int64{2, 1, 4, 3}
	v := NewFromInt(s)
	v.ExtendFromSlice([]interface{}{2, 1, 4, 3})
	s = append(s, s...)
	for index, item := range v.Range() {
		if item != s[index] {
			t.Errorf("%d (%T) != %d (%T)", item, item, s[index], s[index])
		}
	}
}

func TestVec_ExtendFromVec(t *testing.T) {
	s := []int64{2, 1, 4, 3}
	v := NewFromInt(s)
	ov := v.Copy()
	v.ExtendFromVec(ov)
	s = append(s, s...)
	for index, item := range v.Range() {
		if item != s[index] {
			t.Errorf("%d (%T) != %d (%T)", item, item, s[index], s[index])
		}
	}
}

func TestVec_Insert(t *testing.T) {
	s := []int64{2, 1, 4, 3}
	v := NewFromInt(s)
	v.Insert(1, 8)
	if v.AtInt(1) != 8 {
		t.Error("insertion failed")
	}
}
