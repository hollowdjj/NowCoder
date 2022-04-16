package jz

/*
输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。
*/

func printListFromTailToHead(head *ListNode) []int {
	res := make([]int, 0)
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		res = append(res, node.Val)
	}
	dfs(head)
	return res
}
