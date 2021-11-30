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
		if res := ReverseKGroup(utility.SliceToList(v.source), v.k); !utility.EqualSliceInt(res.Slice(), v.wanting) {
			t.Errorf("ReverseKGroup(%q) = %q", v.source, res.Slice())
		}
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := ReverseKGroupAdvanced(head, v.k)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("ReverseKGroupAdvanced(%v,%d)=%v", v.source, v.k, res.Slice())
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
		res := RemoveNthFromEnd(head, v.n).Slice()
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
		if !resList1.Equal(utility.SliceToList(v.sum)) {
			t.Errorf("AddInList(%v,%v) = %v", v.list1, v.list2, resList1.Slice())
		}
		resList2 := AddInListAdvanced(head1, head2)
		if !resList2.Equal(utility.SliceToList(v.sum)) {
			t.Errorf("AddInListAdvanced(%v,%v) = %v", v.list1, v.list2, resList2.Slice())
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
		if res := ReverseBetween(source, v.m, v.n); !res.Equal(wanting) {
			t.Errorf("ReverseBetween(%v,%d,%d) = %v", v.source, v.m, v.n, res.Slice())
		}
	}
}

func TestDeleteDuplicatesII(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := DeleteDuplicatesII(head)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("DeleteDuplicatesII(%v)=%s", v.source, res)
		}
	}
}

func TestOddEvenList(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 3, 5, 2, 4, 6}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		newHead := OddEvenList(head)
		if !newHead.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("OddEvenList(%v)=%s", v.source, newHead)
		}
	}
}

func TestReorderListAdvanced(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		ReorderListAdvanced(head)
		if !head.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("ReorderListAdvanced(%v)=%s", v.source, head)
		}
	}
}

func TestYsf(t *testing.T) {
	data := []struct {
		m, n    int
		wanting int
	}{
		{2, 5, 3},
	}

	for _, v := range data {
		if res := Ysf(v.m, v.n); res != v.wanting {
			t.Errorf("Ysf(%v,%v)=%v", v.m, v.n, res)
		}
	}

	data1 := []struct {
		source  []int
		m       int
		wanting int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 3},
		{[]int{1}, 1, 1},
	}

	for _, v := range data1 {
		head := utility.SliceToList(v.source)
		nHead := head
		for nHead.Next != nil {
			nHead = nHead.Next
		}
		nHead.Next = head
		if res := YsfByList(head, v.m).Val; res != v.wanting {
			t.Errorf("YsfByList(%d,%d)=%d", v.source[len(v.source)-1], v.m, res)
		}
	}
}

func TestPartition(t *testing.T) {
	data := []struct {
		source  []int
		x       int
		wanting []int
	}{
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 1}, 5, []int{1, 2, 3, 4, 1}},
		{[]int{}, 5, []int{}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		newHead := Partition(head, v.x)
		if !newHead.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("Partition(%s,%d)=%s", head, v.x, newHead)
		}
	}
}

func TestSwapLinkedPair(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := SwapLinkedPair(head)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("SwapLinkedPair(%v)=%s", v.source, res)
		}
	}
}

func TestPlusOne(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{9}, []int{1, 0}},
		{[]int{}, []int{}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := PlusOne(head)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("PlusOne(%v)=%s", v.source, res)
		}
	}
}

func TestSortLinkedList(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{1, 3, 2, 2, 3, 1}, []int{1, 1, 2, 2, 3, 3}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := SortLinkedList(head)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("SortLinkedList(%v)=%s", v.source, res)
		}
	}
}

func TestMergeList(t *testing.T) {
	data := []struct {
		source1 []int
		source2 []int
		wanting []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{}, []int{1}},
	}

	for _, v := range data {
		head1 := utility.SliceToList(v.source1)
		head2 := utility.SliceToList(v.source2)
		res := MergeList(head1, head2)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("MergeList(%v,%v)=%s", v.source1, v.source2, res)
		}
	}
}

func TestRotateLinkedList(t *testing.T) {
	data := []struct {
		source  []int
		k       int
		wanting []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{4, 5, 1, 2, 3}},
	}

	for _, v := range data {
		head := utility.SliceToList(v.source)
		res := RotateLinkedList(head, v.k)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("RotateLinkedList(%v,%d)=%s", v.source, v.k, res)
		}
	}
}

func TestMergeKLists(t *testing.T) {
	data := []struct {
		source  [][]int
		wanting []int
	}{
		{[][]int{}, []int{}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, v := range data {
		lists := make([]*utility.ListNode, len(v.source))
		for i := 0; i < len(lists); i++ {
			lists[i] = utility.SliceToList(v.source[i])
		}
		res := MergeKLists(lists)
		if !res.Equal(utility.SliceToList(v.wanting)) {
			t.Errorf("MergeKLists(%v)=%s", v.source, res)
		}
	}
}
