package jz_II

/*
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

func pathSum50(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	res += dfs50(root, 0, targetSum)
	res += pathSum50(root.Left, targetSum)
	res += pathSum50(root.Right, targetSum)
	return res
}

func dfs50(root *TreeNode, sum int, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	sum += root.Val
	if sum == targetSum {
		res++
	}
	res += dfs(root.Left, sum, targetSum)
	res += dfs(root.Right, sum, targetSum)
	return res
}
