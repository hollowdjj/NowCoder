package jz_II

/*
给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值
为 0 的子树。节点 node 的子树为 node 本身，以及所有 node 的后代。
			  1								1
		0		    1       ====>     			1
	0 	    0	1       1					1		1
*/

func pruneTree(root *TreeNode) *TreeNode {
	//后续遍历
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
