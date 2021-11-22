package utility

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(slice []int) *ListNode {
	n := len(slice)
	if n < 1 {
		return nil
	}
	pHead := &ListNode{
		Val:  slice[0],
		Next: nil,
	}
	temp := pHead
	for i := 1; i < n; i++ {
		temp.Next = &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		temp = temp.Next
	}
	return pHead
}

func PrintList(pHead *ListNode) {
	fmt.Print("[")
	for pHead != nil {
		fmt.Printf(" %d ", pHead.Val)
		pHead = pHead.Next
	}
	fmt.Println("]")
}

func ListToSlice(head *ListNode) (res []int) {
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return
}

func EqualSliceInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
