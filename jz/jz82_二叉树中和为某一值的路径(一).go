package jz

/*
给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n
*/

func hasPathSum1(root *TreeNode, sum int) bool {
	res := false
	var dfs func(root *TreeNode, rest int)
	dfs = func(root *TreeNode, rest int) {
		if root == nil || res {
			return
		}
		if root.Left == nil && root.Right == nil && rest == root.Val {
			res = true
			return
		}
		dfs(root.Left, rest-root.Val)
		dfs(root.Right, rest-root.Val)
	}
	dfs(root, sum)
	return res
}
