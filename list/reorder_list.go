package list

/*
描述
将给定的单链表L0->L1->L2->L3......Ln，重排序为L0->Ln->L1->Ln-1....
要求使用原地算法，不能只改变节点内部的值，需要对实际的节点进行交换。

要求：空间复杂度O(n)并在链表上进行操作而不新建链表，时间复杂度O(n)
进阶：空间复杂度O(1),时间复杂度O(n)
*/

//ReorderList 重排链表
func ReorderList(head *ListNode) {
	//首先用快慢指针找到中点
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next

	}

	//断开链表
	if prev == nil {
		return
	}
	prev.Next = nil

	//翻转后半部分链表
	head2 := reverse(slow)

	//交叉合并
	dummy := &ListNode{Val: -1}
	nHead := dummy
	for head != nil && head2 != nil {
		nHead.Next = head
		head = head.Next
		nHead = nHead.Next
		nHead.Next = head2
		head2 = head2.Next
		nHead = nHead.Next
	}
	if head != nil {
		nHead.Next = head
	} else {
		nHead.Next = head2
	}

	head = dummy.Next
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
