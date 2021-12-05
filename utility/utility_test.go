package utility

import (
	"testing"
)

func TestInsertKth(t *testing.T) {
	data := []struct {
		source  []int
		k       int
		t       *ListNode
		wanting []int
	}{
		{[]int{1, 2}, -1, &ListNode{Val: -1}, []int{1, 2}},
		{[]int{1, 2}, 0, &ListNode{Val: -1}, []int{-1, 1, 2}},
		{[]int{1, 2}, 2, &ListNode{Val: -1}, []int{1, 2, -1}},
		{[]int{1, 2}, 3, &ListNode{Val: -1}, []int{1, 2, -1}},
		{[]int{1, 2, 3, 4, 5}, 3, &ListNode{Val: -1}, []int{1, 2, 3, -1, 4, 5}},
	}

	for _, v := range data {
		head := SliceToList(v.source)
		res := InsertKth(head, v.t, v.k)
		if !res.Equal(SliceToList(v.wanting)) {
			t.Errorf("InsertKth(%v,%v,%d)=%s", v.source, v.t, v.k, res)
		}
	}
}

func TestEqualTwoDimSlice(t *testing.T) {
	data := []struct {
		s1, s2  [][]int
		wanting bool
	}{
		{[][]int{{1}, {1, 2}}, [][]int{{1}, {1, 2}, {}}, false},
		{[][]int{{1}, {1, 2}}, [][]int{{1}, {1, 2}}, true},
		{[][]int{{}}, [][]int{{}}, true},
		{[][]int{{}}, [][]int{{}, {}, {1}}, false},
	}

	for _, v := range data {
		if res := EqualTwoDimSlice(v.s1, v.s2); res != v.wanting {
			t.Errorf("EqualTwoDimSlice(%v,%v)=%v", v.s1, v.s2, res)
		}
	}
}
