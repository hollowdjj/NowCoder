package list

import "nowcoder/utility"

/*
给定一个用单链表表示的整数，然后把这个整数加一，并返回加一后的整数的单链表表示。
例如：输入{1,2,3}，返回值{1,2,4}  123 + 1 = 124
      输入{9}，返回值{1,0}        9 + 1 = 10
      输入{1,9,9}，返回值{2,0,0}
*/

func PlusOne(head *utility.ListNode) *utility.ListNode {
	if head == nil {
		return nil
	}
	dummy := &utility.ListNode{Val: 0, Next: head}
	//翻转链表
	head = ReverseList(dummy)
	head.Val += 1
	nHead := head
	//遍历翻转后链表
	for nHead != nil {
		if nHead.Val > 9 {
			nHead.Val = 0
			nHead.Next.Val += 1
		}
		nHead = nHead.Next
	}
	//再次翻转链表
	head = ReverseList(head)
	if head.Val == 0 {
		return head.Next
	}
	return head
}
