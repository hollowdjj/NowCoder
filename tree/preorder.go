package tree

import "nowcoder/utility"

/*
二叉树的前序遍历
*/

func PreOrderTraversal(root *utility.TreeNode) []int {
	res := make([]int, 0)
	var preorder func(root *utility.TreeNode)
	preorder = func(root *utility.TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}
