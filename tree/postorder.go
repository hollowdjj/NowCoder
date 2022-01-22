package tree

import "nowcoder/utility"

/*
二叉树的后序遍历
*/

func PostOrderTraversal(root *utility.TreeNode) []int {
	res := make([]int, 0)
	var postorder func(root *utility.TreeNode)
	postorder = func(root *utility.TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}
	postorder(root)
	return res
}
