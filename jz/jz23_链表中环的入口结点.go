package jz

/*
给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。
*/

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = pHead
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}
