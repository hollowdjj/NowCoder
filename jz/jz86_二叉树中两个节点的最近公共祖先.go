package jz

/*
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

注：本题保证二叉树中每个节点的val值均不相同。
*/

func lowestCommonAncestor2(root *TreeNode, o1 int, o2 int) int {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Val == o1 || root.Val == o2 {
			return root.Val
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left != 0 && right != 0 {
			return root.Val
		}
		if left != 0 {
			return left
		}
		return right
	}
	return dfs(root)
}
