package stack

import "testing"

func TestMaxArea(t *testing.T) {
	data := []struct {
		heights []int
		wanting int
	}{
		{[]int{1, 4, 5, 2}, 8},
	}

	for _, v := range data {
		if res := MaxArea(v.heights); res != v.wanting {
			t.Errorf("MaxArea(%v)=%v", v.heights, res)
		}
	}
}
