package jz_II

/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和是路径中各节点值的总和。

给定一个二叉树的根节点 root ，返回其 最大路径和，即所有路径上节点值之和的最大值。
*/

func maxPathSum(root *TreeNode) int {
	res := root.Val
	var dfs func(root *TreeNode, sum int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		left := max(dfs(root.Left, 0), 0)
		right := max(dfs(root.Right, 0), 0)
		res = max(res, root.Val+left+right)
		return root.Val + max(left, right)
	}
	dfs(root, 0)
	return res
}
