package jz_II

/*
给定两个 非空链表 l1和 l2来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。
将这两数相加会返回一个新的链表。可以假设除了数字0之外，这两个数字都不会以零开头。

输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//翻转两链表
	head1, head2 := reverse(l1), reverse(l2)
	dummy := &ListNode{Val: -1}
	head := dummy
	carry := 0
	for head1 != nil && head2 != nil {
		num := head1.Val + head2.Val + carry
		head.Next = &ListNode{Val: num % 10}
		head = head.Next
		carry = num / 10

		head1 = head1.Next
		head2 = head2.Next
	}

	for head1 != nil {
		num := head1.Val + carry
		head.Next = &ListNode{Val: num % 10}
		head = head.Next
		carry = num / 10
		head1 = head1.Next
	}

	for head2 != nil {
		num := head2.Val + carry
		head.Next = &ListNode{Val: num % 10}
		head = head.Next
		carry = num / 10
		head2 = head2.Next
	}

	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}
	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
