package jz

/*
给定一个单链表的头结点pHead(该头节点是有值的，比如在下图，它的val是1)，长度为n，反转该链表后，返回新链表的表头。
*/

func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	for pHead != nil {
		next := pHead.Next
		pHead.Next = prev
		prev = pHead
		pHead = next
	}
	return prev
}
