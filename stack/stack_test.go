package stack

import "testing"

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

func TestQueue(t *testing.T) {
	Push(2)
	Pop()
	Push(1)
	Pop()
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
