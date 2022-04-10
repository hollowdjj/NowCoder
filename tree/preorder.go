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

//非递归实现二叉树的前序遍历
//              1
//         2         3
//     4       5
func PreOrderByLoop(root *TreeNode) []int {
	//使用栈模拟递归
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			//一直遍历左子树，直到叶子节点。即1-->2-->4
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		//得到叶子节点的父节点，然后处理右子树。例如节点5
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = top.Right
		}
	}
	return res
}
