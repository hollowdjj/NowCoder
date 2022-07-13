package jz_II

/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能
看到的节点值。
*/

func rightSideViewDFS(root *TreeNode) []int {
	//深度优先搜索。使用i == len(res)确保回溯的时候不会添加其他节点
	var res []int
	var dfs func(root *TreeNode, i int)
	dfs = func(root *TreeNode, i int) {
		if root == nil {
			return
		}
		if i == len(res) {
			res = append(res, root.Val)
		}
		dfs(root.Right, i+1)
		dfs(root.Left, i+1)
	}
	dfs(root, 0)
	return res
}

func rightSideViewBFS(root *TreeNode) []int {
	//广度优先搜索，层序遍历
	var res []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:]
			if i == 0 {
				res = append(res, front.Val)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
			if front.Left != nil {
				q = append(q, front.Left)
			}
		}
	}
	return res
}
