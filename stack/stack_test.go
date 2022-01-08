package stack

import "testing"

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
