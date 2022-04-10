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

//非递归实现二叉树的后序遍历
func postOrderByLoop(root *TreeNode) []int {
	//后序遍历的顺序是：左右根。因此， 必须要保证左右子节点都访问过了，才能访问
	//本节点
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := []*TreeNode{root}
	var prev, curr *TreeNode
	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		if curr.Left == nil && curr.Right == nil ||
			prev != nil && (prev == curr.Left || prev == curr.Right) {
			//当节点是叶子节点或者节点的左右子节点都被访问过了时，才能访问本节点
			res = append(res, curr.Val)
			stack = stack[:len(stack)-1]
			prev = curr
		} else {
			//访问左右子树
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
		}
	}
	return res
}
