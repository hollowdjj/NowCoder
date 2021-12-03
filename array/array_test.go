package array

import (
	"nowcoder/utility"
	"testing"
)

func TestMoreThanHalfNum(t *testing.T) {
	data := []struct {
		source  []int
		wanting int
	}{
		{[]int{1, 1, 2}, 1},
		{[]int{1}, 1},
		{[]int{}, 0},
		{[]int{1, 2, 2, 2, 3, 3}, 2},
	}

	for _, v := range data {
		if res := MoreThanHalfNum(v.source); res != v.wanting {
			t.Errorf("MoreThanHalfNum(%v)=%v", v.source, res)
		}
	}
}

func TestReorderArray(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 4, 2, 3, 5}, []int{1, 3, 5, 4, 2}},
	}

	for _, v := range data {
		if res := ReorderArray(v.source); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("ReorderArray(%v)=%v", v.source, res)
		}
		var temp []int
		temp = append(temp, v.source...)
		if res := ReorderArrayAdvanced(v.source); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("ReorderArray(%v)=%v", temp, res)
		}
	}
}
