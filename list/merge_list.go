package list

import "nowcoder/utility"

/*
输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
要求：空间复杂度 O(1)，时间复杂度 O(n)
如输入{1,5,9},{2,3,4,7}时，合并后的链表为{1,2,3,4,5,7,9}
*/

//MergeList 合并两个递增的链表。
func MergeList(head1 *utility.ListNode, head2 *utility.ListNode) *utility.ListNode {
	nHead1, nHead2 := head1, head2
	dummy := &utility.ListNode{Val: -1}
	temp := dummy
	for nHead1 != nil && nHead2 != nil {
		if nHead1.Val <= nHead2.Val {
			temp.Next = nHead1
			nHead1 = nHead1.Next
		} else {
			temp.Next = nHead2
			nHead2 = nHead2.Next
		}
		temp = temp.Next
	}
	if nHead1 != nil {
		temp.Next = nHead1
	} else {
		temp.Next = nHead2
	}

	return dummy.Next
}
