package jz

/*
合并两个升序链表
*/

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	tail := dummy
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			tail.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			tail.Next = pHead2
			pHead2 = pHead2.Next
		}
		tail = tail.Next
	}
	if pHead1 != nil {
		tail.Next = pHead1
	} else {
		tail.Next = pHead2
	}

	return dummy.Next
}
