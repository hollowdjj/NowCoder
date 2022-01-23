package utility

import (
	"fmt"
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

func TestQueue(t *testing.T) {
	q := CreateQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Printf("Pop: %d\t", *q.Pop())
	fmt.Printf("Size: %d\n", q.Size())
	fmt.Printf("Pop: %d\t", *q.Pop())
	fmt.Printf("Size: %d\n", q.Size())
	fmt.Printf("Pop: %d\t", *q.Pop())
	fmt.Printf("Size: %d\n", q.Size())
	fmt.Printf("Pop: %d\t", *q.Pop())
	fmt.Printf("Size: %d\n", q.Size())
}

func TestSliceToBinaryTree(t *testing.T) {
	data := []struct {
		source  []int
		wanting []int
	}{
		{[]int{1, 2, 3, 4, '#', 5, 6, '#', 7, '#', '#', '#', '#', 8, '#'},
			[]int{1, 2, 3, 4, '#', 5, 6, '#', 7, '#', '#', '#', '#', 8, '#'}},
	}

	for _, v := range data {
		root := SliceToBinaryTree(v.source, 0)
		temp := root.LevelOrder()
		if !EqualSliceInt(v.wanting, temp) {
			t.Errorf("SliceToBinaryTree(%v)=%v", v.source, temp)
		}
	}
}

func TestLevelOrder(t *testing.T) {
	nodes := make([]*TreeNode, 5)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = &TreeNode{Val: i + 1}
	}
	nodes[0].Left, nodes[0].Right = nodes[1], nodes[2]
	nodes[1].Left, nodes[2].Right = nodes[3], nodes[4]

	if res := nodes[0].LevelOrder(); !EqualSliceInt(res, []int{1, 2, 3, 4, '#', '#', 5}) {
		t.Errorf("LevelOrder(%v)=%v", nodes[0], res)
	}
}
