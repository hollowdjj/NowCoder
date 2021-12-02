package list

import "nowcoder/utility"

/*
给定链表的头节点，旋转链表，降链表每个节点往右移动 k 个位置，原链表后 k 个位置的节点则依次移动到链表头。
即，例如链表 ： 1->2->3->4->5 k=2 则返回链表 4->5->1->2->3
数据范围：链表中节点数满足 0<=n<=1000    0<=k<=10^9
*/

//RotateLinkedList 旋转链表
func RotateLinkedList(head *utility.ListNode, k int) *utility.ListNode {
	if head == nil {
		return nil
	}
	//先遍历一次链表统计节点个数
	tail, count := head, 1
	for ; tail.Next != nil; tail = tail.Next {
		count++
	}
	//将倒数第index个节点以及其后面的所有节点移动到链表的最前面
	index := k % count
	if index == 0 {
		return head
	}
	slow, fast := head, head
	var prev *utility.ListNode
	for i := 0; i < index; i++ {
		//快指针先走k步
		fast = fast.Next
	}
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = nil
	tail.Next = head

	return slow
}
