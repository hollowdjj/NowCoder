package utility

import (
	"bytes"
	"fmt"
)

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

//SliceToList 将slice转换成链表。若slice为空或者nil时，返回nil
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

//EqualSliceInt 判断两[]int类型的slice是否相等
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

type StackInt []int //定义一个stack类型

func (s *StackInt) Pop() int {
	//先入后出
	if len([]int(*s)) == 0 {
		panic("Empty queue")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return res
}

func (s *StackInt) Push(node int) {
	*s = append(*s, node)
}

func (s *StackInt) Empty() bool {
	if len([]int(*s)) == 0 {
		return true
	}
	return false
}
