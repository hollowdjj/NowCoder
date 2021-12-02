package list

import "nowcoder/utility"

/*
给出一个长度为n的单链表和一个值x，单链表的每一个值为listi，请返回一个链表的头结点，要求新链表
中小于x的节点全部位于大于等于x的节点左侧，并且两个部分之内的节点之间与原来的链表要保持相对顺序不变。
例如：
给出1→4→3→2→5→2和x=3
返回1→2→2→4→3→5
要求：时间复杂度O(n)，空间复杂度O(1)
*/

//Partition 分隔链表
func Partition(head *utility.ListNode, x int) *utility.ListNode {
	smaller := &utility.ListNode{
		Val:  101,
		Next: nil,
	}
	larger := &utility.ListNode{
		Val:  102,
		Next: nil,
	}
	temp1, temp2 := smaller, larger
	for head != nil {
		if head.Val < x {
			temp1.Next = head
			temp1 = temp1.Next
		} else {
			temp2.Next = head
			temp2 = temp2.Next
		}
		head = head.Next
	}
	temp1.Next = larger.Next
	temp2.Next = nil
	return smaller.Next
}
