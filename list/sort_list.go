package list

import (
	"math"
)

/*
给定一个节点数为n的无序单链表，对其按升序排序。
要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
*/

func SortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, mid := findMid(head)
	prev.Next = nil //断链
	return mergeTowList(mergeSort(head), mergeSort(mid))
}

func mergeTowList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	nHead := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			nHead.Next = l1
			l1 = l1.Next
		} else {
			nHead.Next = l2
			l2 = l2.Next
		}
		nHead = nHead.Next
	}

	if l1 != nil {
		nHead.Next = l1
	} else {
		nHead.Next = l2
	}
	return dummy.Next
}

func findMid(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return prev, slow
}

func sortList(head *ListNode) *ListNode {
	//链表的插入排序
	dummy := &ListNode{Val: math.MinInt32}
	tail := dummy

	for head != nil {
		//将head从原链表中断开
		next := head.Next
		head.Next = nil

		if head.Val < tail.Val {
			//插入排序。从dummy后面开始找到第一个大于head.Val的节点，然后将head插入到它前面
			prev := dummy
			nHead := dummy.Next
			for nHead.Val < head.Val {
				prev = nHead
				nHead = nHead.Next
			}

			//将head挂载到prev后面
			prev.Next = head
			head.Next = nHead
		} else {
			tail.Next = head
			tail = tail.Next
		}
		head = next
	}
	return dummy.Next
}
