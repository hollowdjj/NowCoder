package list

import "nowcoder/utility"

/*
给定一个奇数位升序，偶数位降序的链表，返回对其排序后的链表。
题面解释：例如链表 1->3->2->2->3->1 是奇数位升序偶数位降序的链表，而 1->3->2->2->3->2 则不符合题目要求。
输入：
[1,3,2,2,3,1]
返回值：
{1,1,2,2,3,3}
*/

func SortLinkedList(head *utility.ListNode) *utility.ListNode {
	//链表奇偶拆分
	odd, even := PartitionList(head)
	//翻转偶链表
	even = ReverseList(even)
	//排序两已排序链表(归并)
	return MergeList(odd, even)
}

//PartitionList 拆分链表，使得奇数位的元素在一个链表，偶数位的元素在另一个链表中
func PartitionList(head *utility.ListNode) (*utility.ListNode, *utility.ListNode) {
	if head == nil {
		return nil, nil
	}
	evenHead := head.Next
	odd, even := head, evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = nil
	return head, evenHead
}
