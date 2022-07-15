package jz_II

/*
给定一个二叉树的根节点root，树中每个节点都存放有一个0到9之间的数字。每条从根节点到叶节点的
路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和。叶节点 是指没有子节点的节点。

		1
	2		3     ---> 数字总和：12+13=25
*/

func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode, num int)
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		num = num*10 + root.Val
		if root.Left == nil && root.Right == nil {
			res += num
			return
		}
		dfs(root.Left, num)
		dfs(root.Right, num)
	}
	dfs(root, 0)
	return res
}
