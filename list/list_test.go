package list

import (
	"nowcoder/utility"
	"testing"
)

func TestIsPail(t *testing.T) {
	data := []struct {
		source  []int
		wanting bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 3}, false},
	}
	for _, v := range data {
		if res := isPailAdvanced(utility.SliceToList(v.source)); res != v.wanting {
			t.Errorf("isPailAdvanced(%v) = %v", v.source, res)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	data := []struct {
		source  []int
		k       int
		wanting []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{}, 2, []int{}},
		{[]int{1, 2}, 1, []int{1, 2}},
	}

	for _, v := range data {
		if res := ReverseKGroup(utility.SliceToList(v.source), v.k); !utility.EqualSliceInt(utility.ListToSlice(res), v.wanting) {
			t.Errorf("ReverseKGroup(%q) = %q", v.source, utility.ListToSlice(res))
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	data := []struct {
		source  []int
		n       int
		wanting []int
	}{
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := utility.ListToSlice(RemoveNthFromEnd(head, v.n))
		if !utility.EqualSliceInt(res, v.wanting) {
			t.Errorf("RemoveNthFromEnd(%v) = %v", v.source, res)
		}
	}
}
