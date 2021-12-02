package list

import (
	"nowcoder/utility"
)

/*
两两交换链表的节点，例如1->2->3 变为2->1->3
*/

//SwapLinkedPair 两两交换链表的节点
func SwapLinkedPair(head *utility.ListNode) *utility.ListNode {
	dummy := &utility.ListNode{Val: -1, Next: head}
	prev, left := dummy, head

	for left != nil && left.Next != nil {
		right := left.Next
		//交换left和right
		next := right.Next
		prev.Next = right
		right.Next = left
		left.Next = next

		prev = left
		left = left.Next
	}

	return dummy.Next
}
