package jz

/*
操作给定的二叉树，将其变换为源二叉树的镜像
*/

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}
