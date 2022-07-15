package jz_II

/*
给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
*/

func convertBST(root *TreeNode) *TreeNode {
	prevSum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		prevSum += root.Val
		root.Val = prevSum
		dfs(root.Left)
	}
	dfs(root)
	return root
}
