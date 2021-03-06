package array

import (
	"fmt"
	"nowcoder/utility"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	BinarySearch([]int{1, 3, 6, 7, 8, 10, 11}, 9)
}

func TestCombinationSum2(t *testing.T) {
	data := []struct {
		source  []int
		target  int
		wanting [][]int
	}{
		//{[]int{100, 10, 20, 70, 60, 10, 50}, 80, [][]int{{10, 10, 60}, {10, 20, 50}, {10, 70}, {20, 60}}},
		{[]int{22, 49, 5, 24, 26}, 77, [][]int{{5, 22, 24, 26}}},
	}

	for _, v := range data {
		if res := CombinationSum2(v.source, v.target); !utility.EqualTwoDimSlice(v.wanting, res) {
			t.Errorf("CombinationSum2(%v,%v)=%v", v.source, v.target, res)
		}
	}
}

func TestCombination3(t *testing.T) {
	data := []struct {
		k, n    int
		wanting [][]int
	}{
		{3, 7, [][]int{{1, 2, 4}}},
	}

	for _, v := range data {
		if res := Combination3(v.k, v.n); !utility.EqualTwoDimSlice(v.wanting, res) {
			t.Errorf("Combination3(%v,%v)=%v", v.k, v.n, res)
		}
	}
}

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

func TestFind(t *testing.T) {
	data := []struct {
		source  [][]int
		target  int
		wanting bool
	}{
		{[][]int{{}}, 1, false},
		{[][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}, 7, true},
		{[][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}, 3, false},
	}

	for _, v := range data {
		if res := Find(v.target, v.source); res != v.wanting {
			t.Errorf("Find(%v,%v)=%v", v.target, v.source, res)
		}
	}
}

func TestFindPeakElement(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{2, 4, 1, 2, 7, 8, 4}, []int{1, 5}},
		{[]int{1, 2, 3, 1}, []int{2}},
	}

	for _, v := range data {
		res := FindPeakElement(v.source)
		res1 := FindPeakElementAdvanced(v.source)
		flag, flag1 := false, false
		for _, w := range v.wanting {
			if w == res {
				flag = true
				break
			}
		}
		for _, w := range v.wanting {
			if w == res1 {
				flag1 = true
				break
			}
		}
		if !flag {
			t.Errorf("FindPeakElement(%v)=%v", v.source, res)
		}
		if !flag1 {
			t.Errorf("FindPeakElementAdvanced(%v)=%v", v.source, res)
		}

	}
}

func TestFlipChess(t *testing.T) {
	data := []struct {
		A       [][]int
		f       [][]int
		wanting [][]int
	}{
		{[][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}}, [][]int{{2, 2}, {3, 3}, {4, 4}}, [][]int{{0, 1, 1, 1}, {0, 0, 1, 0}, {0, 1, 1, 0},
			{0, 0, 1, 0}}},
	}

	for _, v := range data {
		if res := FlipChess(v.A, v.f); !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("FlipChess(%v,%v)=%v", v.A, v.f, res)
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

func TestInsertInterval(t *testing.T) {
	data := []struct {
		source  [][2]int
		add     [][2]int
		wanting [][2]int
	}{
		{[][2]int{{2, 5}, {6, 11}}, [][2]int{{5, 6}}, [][2]int{{2, 11}}},
		{[][2]int{{1, 2}, {6, 11}}, [][2]int{{3, 4}}, [][2]int{{1, 2}, {3, 4}, {6, 11}}},
		{[][2]int{{1, 3}, {6, 9}}, [][2]int{{2, 5}}, [][2]int{{1, 5}, {6, 9}}},
		{[][2]int{{1, 2}, {3, 4}}, [][2]int{{0, 5}}, [][2]int{{0, 5}}},
	}

	for _, v := range data {
		intervals := SliceToInterval(v.source)
		add := SliceToInterval(v.add)[0]
		wanting := SliceToInterval(v.wanting)
		if res := InsertInterval(intervals, add); !EqualIntervalSlice(res, wanting) {
			t.Errorf("InsertInterval(%v,%v)=%v", v.source, v.add, res)
		}
	}
}

func TestInversePairs(t *testing.T) {
	data := []struct {
		source  []int
		wanting int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 0}, 7},
	}

	for _, v := range data {
		if res := InversePairs(v.source); res != v.wanting {
			t.Errorf("InversePairs(%v)=%v", v.source, res)
		}
	}
}

func TestMatrixProduct(t *testing.T) {
	data := []struct {
		a       [][]int
		b       [][]int
		wanting [][]int
	}{
		{[][]int{{1, 2}, {3, 2}}, [][]int{{3, 4}, {2, 1}}, [][]int{{7, 6}, {13, 14}}},
	}

	for _, v := range data {
		if res := MatrixProduct(v.a, v.b); !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("MatrixProduct(%v,%v)=%v", v.a, v.b, res)
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
		if res := MaxProductOfThreeNum(v.source); res != v.wanting {
			t.Errorf("MaxProductOfThreeNum(%v)=%v", v.source, res)
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

		if res := RotateMatrixAdvanced(v.source, len(v.source)); !utility.EqualTwoDimSlice(res, v.wanting) {
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

func TestSubArrayMaxProduct(t *testing.T) {
	data := []struct {
		source  []float64
		wanting float64
	}{
		{[]float64{-2.5, 4, 0, 3, 0.5, 8, -1}, 12.0},
	}

	for _, v := range data {
		if res := SubArrayMaxProduct(v.source); res != v.wanting {
			t.Errorf("SubArrayMaxProduct(%v)=%v", v.source, res)
		}
	}
}

func TestSubSets(t *testing.T) {
	data := []struct {
		source  []int
		wanting [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, v := range data {
		if res := SubSets(v.source); !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("SubSets(%v)=%v", v.source, res)
		}
	}
}

func TestSubsets2(t *testing.T) {
	data := []struct {
		source  []int
		wanting [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	}

	for _, v := range data {
		if res := Subsets2(v.source); !utility.EqualTwoDimSlice(v.wanting, res) {
			t.Errorf("Subsets2(%v)=%v", v.source, res)
		}
	}
}

func TestTemperatures(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{1, 2, 3}, []int{1, 1, 0}},
		{[]int{1}, []int{0}},
	}

	for _, v := range data {
		if res := Temperatures(v.source); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("Temperatures(%v)=%v", v.source, res)
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

func TestSort(t *testing.T) {
	data := []struct {
		arr     []int
		wanting []int
	}{
		{[]int{3, 2, 4, 1, 5}, []int{1, 2, 3, 4, 5}},
		//{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		//{[]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}},
		//{[]int{3, 2, 4, 5, 1, 5, 6, 12}, []int{1, 2, 3, 4, 5, 5, 6, 12}},
	}

	for _, v := range data {
		temp := make([]int, 0)
		temp = append(temp, v.arr...)
		if res := MergeSort(v.arr); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("MergeSort(%v)=%v", temp, res)
		}
		if res := HeapSort(v.arr); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("HeapSort(%v)=%v", temp, res)
		}
		if res := HeapSort(v.arr); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("HeapSort(%v)=%v", temp, res)
		}
	}
}

func TestPermuteUnique(t *testing.T) {
	data := []struct {
		arr     []int
		wanting [][]int
	}{
		{[]int{2, 2, -1}, [][]int{{-1, 2, 2}, {2, -1, 2}, {2, 2, -1}}},
	}

	for _, v := range data {
		if res := PermuteUnique(v.arr); !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("PermutaUnique(%v)=%v", v.arr, res)
		}
	}
}

func TestMedianNum(t *testing.T) {
	arr := []int{383, 886, 777, 915, 793, 335, 386, 492}
	for _, v := range arr {
		Insert(v)
		fmt.Println(GetMedian())
	}
}
