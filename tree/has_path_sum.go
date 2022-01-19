package tree

import "nowcoder/utility"

/*
二叉树中和为某一值的路径(一)
给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n
*/

func HasPathSum(root *utility.TreeNode, sum int) bool {
	//很容易看出来这道题的思路是dfs。唯一需要注意的是，题目中要求路径的终点
	//必须是叶子节点
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		//注意这里的判断条件。root.Val == sum判断路径和是否正确。
		//root.Left == nil && root.Right == nil判断当前节点是不是叶子节点
		return true
	}
	return HasPathSum(root.Left, sum-root.Val) || HasPathSum(root.Right, sum-root.Val)
}
