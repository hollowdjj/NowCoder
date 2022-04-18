package jz

import (
	"bytes"
	"fmt"
)

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Equal(another *ListNode) bool {
	for head != nil && another != nil {
		if head.Val != another.Val {
			return false
		}
		head, another = head.Next, another.Next
	}
	if head != nil || another != nil {
		return false
	}

	return true
}

func (head *ListNode) String() string {
	str := bytes.Buffer{}
	str.WriteRune('[')
	for head != nil {
		str.WriteString(fmt.Sprintf(" %d ", head.Val))
		head = head.Next
	}
	str.WriteRune(']')
	return str.String()
}

func (head *ListNode) Slice() []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

//TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}
