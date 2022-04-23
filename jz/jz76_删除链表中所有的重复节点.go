package jz

/*
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
例如，链表 1->2->3->3->4->4->5  处理后为 1->2->5
*/

func deleteDuplication(pHead *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: pHead}
	prev, curr := dummy, pHead
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			//找到下一个不重复的点
			temp := curr
			for temp != nil && temp.Val == curr.Val {
				temp = temp.Next
			}
			prev.Next = temp
			curr = temp
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return dummy.Next
}
