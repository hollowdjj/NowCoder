package jz

/*
输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度
为树的深度，根节点的深度视为1
*/

func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	return max(TreeDepth(pRoot.Left), TreeDepth(pRoot.Right)) + 1
}
