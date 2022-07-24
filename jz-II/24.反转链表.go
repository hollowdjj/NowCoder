package jz_II

func reverseList(head *ListNode) *ListNode {
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
