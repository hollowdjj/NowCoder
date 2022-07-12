package jz_II

/*
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
*/

func findBottomLeftValueBFS(root *TreeNode) int {
	//广度优先搜索
	q := []*TreeNode{root}
	res := 0
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:]
			if i == 0 {
				res = front.Val
			}

			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}
	}
	return res
}

func findBottomLeftValueDFS(root *TreeNode) int {
	//前序遍历
	res, maxHight := 0, 0
	var dfs func(root *TreeNode, height int)
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		if height > maxHight {
			res = root.Val
			maxHight = height
		}
		dfs(root.Left, height+1)
		dfs(root.Right, height+1)
	}
	return res
}
