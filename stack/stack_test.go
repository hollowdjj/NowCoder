package stack

import (
	"nowcoder/utility"
	"testing"
)

func TestCalculate1(t *testing.T) {
	data := []struct {
		s       string
		wanting int
	}{
		{"(1+2)+1", 4},
		{"-1", -1},
	}

	for _, v := range data {
		if res := Calculate1(v.s); res != v.wanting {
			t.Errorf("Calculate1(%v)=%v", v.s, res)
		}
	}
}

func TestCalculate2(t *testing.T) {
	data := []struct {
		s       string
		wanting int
	}{
		{"(1+2)+1", 4},
		{"-1", -1},
		{"(2*3+1)+(1*2-1)", 8},
	}

	for _, v := range data {
		if res := Calculate2(v.s); res != v.wanting {
			t.Errorf("Calculate1(%v)=%v", v.s, res)
		}
	}
}

func TestIsValid(t *testing.T) {
	data := []struct {
		s       string
		wanting bool
	}{
		{"]", false},
		{"[]", true},
		{"[)", false},
		{"{([]){", false},
		{"([]){}", true},
		{"[](){{", false},
	}

	for _, v := range data {
		if res := IsValid(v.s); res != v.wanting {
			t.Errorf("IsValid(%v)=%v", v.s, res)
		}
	}
}

func TestMaxArea(t *testing.T) {
	data := []struct {
		heights []int
		wanting int
	}{
		{[]int{1, 4, 5, 2}, 8},
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{3, 2, 5, 6, 1, 4, 4}, 10},
	}

	for _, v := range data {
		if res := MaxArea(v.heights); res != v.wanting {
			t.Errorf("MaxArea(%v)=%v", v.heights, res)
		}
	}

	for _, v := range data {
		if res := MaxAreaMonotoneStack(v.heights); res != v.wanting {
			t.Errorf("MaxAreaMonotoneStack(%v)=%v", v.heights, res)
		}
	}
}

func TestMaximalRectangle(t *testing.T) {
	data := []struct {
		matrix  [][]int
		wanting int
	}{
		{[][]int{{1}}, 1},
		{[][]int{{1, 1}, {0, 1}}, 2},
		{[][]int{{0, 0}, {0, 0}}, 0},
		{[][]int{{1, 0, 1, 0, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}, 8},
	}

	for _, v := range data {
		if res := MaximalRectangle(v.matrix); res != v.wanting {
			t.Errorf("MaximalRectangle(%v)=%v", v.matrix, res)
		}
	}
}

func TestPrint(t *testing.T) {
	data := []struct {
		s       []int
		wanting [][]int
	}{
		{[]int{1, 2, 3, '#', '#', 4, 5}, [][]int{{1}, {3, 2}, {4, 5}}},
		{[]int{1, 2, 3, '#', '#', 4, 5, '#', '#', '#', '#', 6, '#', '#', '#'}, [][]int{{1}, {3, 2}, {4, 5}, {6}}},
	}

	for _, v := range data {
		root := utility.SliceToBinaryTree(v.s, 0)
		res := Print(root)
		if !utility.EqualTwoDimSlice(res, v.wanting) {
			t.Errorf("Print(%v)=%v", v.s, res)
		}
	}

}

func TestQueue(t *testing.T) {
	Push(2)
	Pop()
	Push(1)
	Pop()
}

func TestOrder(t *testing.T) {
	data := []struct {
		a       []int
		wanting []int
	}{
		{[]int{2, 1, 5, 3, 4}, []int{5, 4, 3, 1, 2}},
	}

	for _, v := range data {
		if res := Order(v.a); !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("Order(%v)=%v", v.a, res)
		}
	}
}

func TestSolve(t *testing.T) {
	data := []struct {
		s       string
		wanting int
	}{
		{"1+2", 3},
		{"(2 + 3 * 4) + 4", 18},
		{"-1", -1},
	}

	for _, v := range data {
		if res := Solve(v.s); res != v.wanting {
			t.Errorf("Solve(%v)=%v", v.s, res)
		}
	}

}
