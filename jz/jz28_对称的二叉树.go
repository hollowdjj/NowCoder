package jz

/*
给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）
*/

func isSymmetrical(pRoot *TreeNode) bool {
	return judge(pRoot, pRoot)
}

func judge(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && judge(root1.Left, root2.Right) && judge(root1.Right, root2.Left)
}
