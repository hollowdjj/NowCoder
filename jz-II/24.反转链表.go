package jz_II

/*
给定链表头节点，翻转链表。
输入：1 2 3 4 5
输出：5 4 3 2 1
*/

func reverseList(head *ListNode) *ListNode {
	//迭代法翻转链表
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseList1(head *ListNode) *ListNode {
	//递归法翻转链表
	var dfs func(prev, curr *ListNode) *ListNode
	dfs = func(prev, curr *ListNode) *ListNode {
		if curr == nil {
			return prev
		}
		next := curr.Next
		curr.Next = prev
		return dfs(curr, next)
	}
	return dfs(nil, head)
}
