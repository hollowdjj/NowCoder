package jz

/*
输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。
*/

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	//快慢指针，快指针先走k步，然后两指针一起走
	slow, fast := pHead, pHead
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
