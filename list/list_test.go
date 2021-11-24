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
		{[]int{1, 2, 3}, 3, []int{3, 2, 1}},
	}

	for _, v := range data {
		if res := ReverseKGroup(utility.SliceToList(v.source), v.k); !utility.EqualSliceInt(utility.ListToSlice(res), v.wanting) {
			t.Errorf("ReverseKGroup(%q) = %q", v.source, utility.ListToSlice(res))
		}
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := ReverseKGroupAdvanced(head, v.k)
		if !utility.EqualList(res, utility.SliceToList(v.wanting)) {
			t.Errorf("ReverseKGroupAdvanced(%v,%d)=%v", v.source, v.k, utility.ListToSlice(res))
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

func TestAddInList(t *testing.T) {
	data := []struct {
		list1 []int
		list2 []int
		sum   []int
	}{
		{[]int{9, 3, 7}, []int{6, 3}, []int{1, 0, 0, 0}},
		{[]int{0}, []int{6, 3}, []int{6, 3}},
	}

	for _, v := range data {
		head1 := utility.SliceToList(v.list1)
		head2 := utility.SliceToList(v.list2)
		resList1 := AddInList(head1, head2)
		if res := utility.EqualList(resList1, utility.SliceToList(v.sum)); !res {
			t.Errorf("AddInList(%v,%v) = %v", v.list1, v.list2, utility.ListToSlice(resList1))
		}
		resList2 := AddInListAdvanced(head1, head2)
		if res := utility.EqualList(resList2, utility.SliceToList(v.sum)); !res {
			t.Errorf("AddInListAdvanced(%v,%v) = %v", v.list1, v.list2, utility.ListToSlice(resList2))
		}
	}
}

func TestReverseBetween(t *testing.T) {
	data := []struct {
		source  []int
		m       int
		n       int
		wanting []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, 1, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, 1, []int{1, 2, 3, 4, 5}},
	}

	for _, v := range data {
		source := utility.SliceToList(v.source)
		wanting := utility.SliceToList(v.wanting)
		if res := ReverseBetween(source, v.m, v.n); !utility.EqualList(res, wanting) {
			t.Errorf("ReverseBetween(%v,%d,%d) = %v", v.source, v.m, v.n, utility.ListToSlice(res))
		}
	}
}
