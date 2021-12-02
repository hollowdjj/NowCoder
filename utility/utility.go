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

//InsertKth 将target插入到以head为头结点的链表的第k个元素后面，如链表长度小于等于k
//则插入到最后面。若k等于0，则表示插入到最前面，k小于0则不会发生任何插入
func InsertKth(head, target *ListNode, k int) *ListNode {
	if k < 0 || head == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head}
	count, nHead := 0, dummy
	for nHead.Next != nil {
		if count == k {
			next := nHead.Next
			nHead.Next = target
			target.Next = next
			return dummy.Next
		}
		nHead = nHead.Next
		count++
	}

	nHead.Next = target
	target.Next = nil
	return dummy.Next
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

//SliceToLoopList 将链表的最后一个元素与k个元素相连形成一个有环链表。k小于1时无环
func SliceToLoopList(slice []int, k int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	//k小于1时无环
	if k < 1 {
		return SliceToList(slice)
	}

	var nodes []*ListNode
	head := &ListNode{Val: slice[0]}
	temp := head
	nodes = append(nodes, head)
	for i := 1; i < len(slice); i++ {
		newNode := &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		nodes = append(nodes, newNode)
		temp.Next = newNode
		temp = temp.Next
	}
	nodes[len(nodes)-1].Next = nodes[k-1]
	return head
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
