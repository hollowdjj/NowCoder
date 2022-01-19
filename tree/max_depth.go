package tree

import "nowcoder/utility"

/*
求给定二叉树的最大深度，
深度是指树的根节点到任一叶子节点路径上节点的数量。
最大深度是所有叶子节点的深度的最大值。
（注：叶子节点是指没有子节点的节点。）
*/

func MaxDepth(root *utility.TreeNode) int {
	//深度优先搜索求解。height表示的是node的深度
	if root == nil {
		return 0
	}
	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
