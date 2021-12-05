package array

import (
	"nowcoder/utility"
	"testing"
)

func TestGetNumberOfK(t *testing.T) {
	data := []struct {
		source  []int
		k       int
		wanting int
	}{
		{[]int{}, 1, 0},
		{[]int{1}, 2, 0},
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 2, 3, 3, 3, 4, 4, 5}, 3, 3},
	}

	for _, v := range data {
		if res := GetNumberOfK(v.source, v.k); res != v.wanting {
			t.Errorf("GetNumberOfK(%v)=%v", v.source, res)
		}
	}
}

func TestMaxProduct(t *testing.T) {
	data := []struct {
		source  []int
		wanting int64
	}{
		{[]int{3, 4, 1, 2}, 24},
		{[]int{-1, -1, 1, 2}, 2},
		{[]int{-2, -2, 1, 1}, 4},
	}

	for _, v := range data {
		if res := MaxProduct(v.source); res != v.wanting {
			t.Errorf("MaxProduct(%v)=%v", v.source, res)
		}
	}
}

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

func TestYangHuiTriangle(t *testing.T) {
	data := []struct {
		num     int
		wanting [][]int
	}{
		{1, [][]int{{1}}},
		{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	}

	for _, v := range data {
		if res := YangHuiTriangle(v.num); !utility.EqualTwoDimSlice(v.wanting, res) {
			t.Errorf("YangHuiTriangle(%d)=%v", v.num, res)
		}
	}
}
