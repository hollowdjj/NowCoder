package list

import "nowcoder/utility"

/*
描述
将给定的单链表L0->L1->L2->L3......Ln，重排序为L0->Ln->L1->Ln-1....
要求使用原地算法，不能只改变节点内部的值，需要对实际的节点进行交换。

要求：空间复杂度O(n)并在链表上进行操作而不新建链表，时间复杂度O(n)
进阶：空间复杂度O(1),时间复杂度O(n)
*/

func ReorderList(head *utility.ListNode) {

}

func ReorderListAdvanced(head *utility.ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	//找到链表的中点
	var prev *utility.ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	prev.Next = nil
	prev = nil
	//将后半部分的链表[slow,tail]翻转，翻转后的头节点为prev
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	//交叉合并两链表
	dummy := &utility.ListNode{
		Val:  -1,
		Next: nil,
	}
	nHead := dummy
	for head != nil && prev != nil {
		temp1 := head.Next
		temp2 := prev.Next
		nHead.Next = head
		nHead = nHead.Next
		nHead.Next = prev
		nHead = nHead.Next
		head, prev = temp1, temp2
	}
	if head != nil {
		nHead.Next = head
	} else {
		nHead.Next = prev
	}
	head = dummy.Next
}
