package jz_II

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：L0→L1→ … → Ln-1→ Ln
请将其重新排列后变为：L0→Ln→L1→Ln-1→L2→Ln-2→ …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

输入: head = [1,2,3,4]
输出: [1,4,2,3]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	//快慢指针找到链表中点
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//翻转后半部分的链表
	prev.Next = nil
	head1 := reverse(slow)
	//交替合并成新链表
	dummy := &ListNode{Val: -1}
	nHead := dummy
	for head != nil && head1 != nil {
		nHead.Next = head
		head = head.Next
		nHead = nHead.Next
		nHead.Next = head1
		head1 = head1.Next
		nHead = nHead.Next
	}
	if head != nil {
		nHead.Next = head
	}
	if head1 != nil {
		nHead.Next = head1
	}
}
