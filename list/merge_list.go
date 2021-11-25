package list

import "nowcoder/utility"

/*
输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
要求：空间复杂度 O(1)，时间复杂度 O(n)
如输入{1,5,9},{2,3,4,7}时，合并后的链表为{1,2,3,4,5,7,9}
*/

func MergeList(pHead1 *utility.ListNode, pHead2 *utility.ListNode) *utility.ListNode {
	start := &utility.ListNode{
		Val:  0,
		Next: nil,
	}
	temp := start //哨兵节点
	for pHead1 == nil && pHead2 == nil {
		if pHead1 != nil && (pHead2 == nil || pHead1.Val <= pHead2.Val) {
			temp.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			temp.Next = pHead2
			pHead2 = pHead2.Next
		}
		temp = temp.Next
	}

	return start.Next
}
