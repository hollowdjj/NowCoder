package easy

//给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead
	for {
		//快指针一次走两步，慢指针一次走一步，最终会在环中的某一点相遇
		if fast == nil {
			return nil
		}

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next
		if slow == fast {
			/*
				快慢指针相遇后，从链表头和相遇点每次移动一格，再次相遇的点即为环的入口。
				证明：设链表头到环入口的距离为a，设链表头到快慢指针第一次相遇点的距离为b
				设该相遇点到环入口的距离为c。则第一次相遇时：
				快指针走过的路程为：a + k(b+c) + b (k>0，即快指针至少需要先走一圈环，才能追上慢指针)
				慢指针走过的路程为：a + b
				而快指针走过的路程一定是慢指针的两倍，故有：
				a + k(b+c) +b = 2a + 2b ----> a + b = kb + kc ----> a = (k-1)(b+c) + c
				即链表头到环入口的距离=相遇点到环入口的距离+（k-1）圈环长度
			*/
			slow = pHead
			for {
				if slow == fast {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
}
