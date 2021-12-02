package list

import "nowcoder/utility"

/*
给定一个链表，删除链表的倒数第n个节点并返回链表的头指针
例如，给出的链表为: 1→2→3→4→5,n=2，删除了链表的倒数第n个节点之后,链表变为1→2→3→5.
备注: 题目保证n一定是有效的
*/

//RemoveNthFromEnd 删除链表倒数第n个节点
func RemoveNthFromEnd(head *utility.ListNode, n int) *utility.ListNode {
	if n == 0 {
		return head
	}

	slow, fast := head, head
	//快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//然后快慢指针一起一次都走一步，循环结束后，slow为倒数第n个节点
	var prev *utility.ListNode
	for fast != nil {
		prev = slow
		slow, fast = slow.Next, fast.Next
	}
	//删除slow指向的节点
	if prev == nil {
		//说明是要删除第一个节点
		head = head.Next
	} else {
		prev.Next = slow.Next
	}

	return head
}
