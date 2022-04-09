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
func PreOrderByLoop(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//使用栈模拟递归
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	res := make([]int, 0)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		//栈是先入后出，所以前序遍历的话要先push右子节点
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}
