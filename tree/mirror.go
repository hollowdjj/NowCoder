package tree

import "nowcoder/utility"

/*
镜像二叉树
操作给定的二叉树，将其装换成镜像的二叉树
*/

func Mirror(root *utility.TreeNode) *utility.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	Mirror(root.Left)
	Mirror(root.Right)
	return root
}
