package array

import (
	"nowcoder/utility"
	"testing"
)

func TestFibonacci(t *testing.T) {
	data := []struct {
		n, wanting int
	}{
		{1, 1},
		{2, 1},
		{4, 3},
	}

	for _, v := range data {
		if res := Fibonacci(v.n); res != v.wanting {
			t.Errorf("Fibonacci(%d)=%d", v.n, res)
		}
	}
}

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

func TestMaxLength(t *testing.T) {
	data := []struct {
		source  []int
		wanting int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{2, 2, 3, 4, 3}, 3},
	}

	for _, v := range data {
		if res := MaxLength(v.source); res != v.wanting {
			t.Errorf("MaxLength(%v)=%d", v.source, res)
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

func TestFindMedianInTwoSortedArray(t *testing.T) {
	data := []struct {
		arr1, arr2 []int
		wanting    int
	}{
		{[]int{1, 2, 3, 4}, []int{3, 4, 5, 6}, 3},
		{[]int{1, 4, 6, 7, 9}, []int{2, 3, 4, 5, 8}, 4},
		{[]int{1}, []int{2}, 1},
	}

	for _, v := range data {
		if res := FindMedianInTwoSortedArray(v.arr1, v.arr2); res != v.wanting {
			t.Errorf("FindMedianInTwoSortedArray(%v,%v)=%d", v.arr1, v.arr2, res)
		}
	}
}

func TestMergeIntervals(t *testing.T) {
	data := []struct {
		source  [][2]int
		wanting [][2]int
	}{
		//{[][2]int{{10, 30}, {80, 100}, {150, 180}, {20, 60}}, [][2]int{{10, 60}, {80, 100}, {150, 180}}},
		//{[][2]int{{0, 10}, {10, 20}}, [][2]int{{0, 20}}},
		//{[][2]int{}, [][2]int{}},
		{[][2]int{{1, 4}, {2, 3}}, [][2]int{{1, 4}}},
	}

	for _, v := range data {
		intervals := SliceToInterval(v.source)
		res := MergeIntervals(intervals)
		wanting := SliceToInterval(v.wanting)
		if !EqualIntervalSlice(wanting, res) {
			t.Errorf("MergeIntervals(%s)=%s", PrintIntervals(intervals), PrintIntervals(res))
		}
	}
}

func TestMinNumberDisappeared(t *testing.T) {
	data := []struct {
		source  []int
		wanting int
	}{
		{[]int{1, 0, 2}, 3},
		{[]int{-2, 3, 4, 1, 5}, 2},
		{[]int{4, 5, 6, 8, 9}, 1},
		{[]int{3, 2, 1}, 4},
	}

	for _, v := range data {
		var array []int
		array = append(array, v.source...)
		if res := MinNumberDisappeared(array); res != v.wanting {
			t.Errorf("MinNumberDisappeared(%v)=%v", v.source, res)
		}
	}
}

func TestMinPathSum(t *testing.T) {
	data := []struct {
		source  [][]int
		wanting int
	}{
		{[][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}}, 12},
		{[][]int{{1, 2, 3}, {1, 2, 3}}, 7},
	}

	for _, v := range data {
		if res := MinPathSum(v.source); res != v.wanting {
			t.Errorf("MinPathSum(%v)=%v", v.source, res)
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

func TestReconstructBinaryTree(t *testing.T) {
	//data := []struct {
	//	pre, vin []int
	//	tree     string
	//}{
	//	{[]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}, "1,2,3,4,#,5,6,#,7,#,#,8"},
	//}
	//
	//for _, v := range data {
	//
	//}
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

func TestReverseArray(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, v := range data {
		var temp []int
		temp = append(temp, v.source...)
		if ReverseArray(v.source); !utility.EqualSliceInt(v.source, v.wanting) {
			t.Errorf("ReverseArray(%v)=%v", temp, v.source)
		}
	}
}

func TestRotateArray(t *testing.T) {
	data := []struct {
		source  []int
		m       int
		wanting []int
	}{
		{[]int{}, 10, []int{}},
		{[]int{1}, 10, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{5, 6, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 8, []int{5, 6, 1, 2, 3, 4}},
	}

	for _, v := range data {
		if res := RotateArray(v.source, v.m, len(v.source)); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("RotateArray(%v,%d,%d)=%v", v.source, v.m, len(v.source), res)
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	data := []struct {
		source  [][]int
		wanting [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}

	for _, v := range data {
		if res := RotateMatrix(v.source, len(v.source)); !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("RotateMatrix(%v,%v)=%v", v.source, len(v.source), res)
		}
	}
}

func TestSpiralOrder(t *testing.T) {
	data := []struct {
		source  [][]int
		wanting []int
	}{
		{[][]int{}, []int{}},
		{[][]int{{2, 3}}, []int{2, 3}},
		{[][]int{{2}, {3}}, []int{2, 3}},
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}

	for _, v := range data {
		if res := SpiralOrder(v.source); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("SpiralOrder(%v)=%v", v.source, res)
		}
	}
}

func TestThreeSum(t *testing.T) {
	data := []struct {
		source  []int
		wanting [][]int
	}{
		{[]int{-2, -2, -2, -2}, [][]int{}},
		{[]int{0}, [][]int{}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{[]int{-10, 0, 10, 20, -10, -40}, [][]int{{-10, -10, 20}, {-10, 0, 10}}},
	}

	for _, v := range data {
		if res := ThreeSum(v.source); !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("ThreeSum(%v)=%v", v.source, res)
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

/*
工具函数测试
*/
