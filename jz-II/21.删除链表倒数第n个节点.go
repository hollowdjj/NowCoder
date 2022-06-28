package jz_II

/*
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//快慢指针，快指针先走n步，然后一起走，当快指针走到尾时，慢指针指向的就是倒数
	//第n个节点
	var prev *ListNode
	slow, fast := head, head
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}
	if prev == nil {
		return head.Next
	}
	prev.Next = slow.Next
	return head
}
