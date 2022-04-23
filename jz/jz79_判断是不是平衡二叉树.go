package jz

/*
输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左
右两个子树都是一棵平衡二叉树。
注：我们约定空树是平衡二叉树
*/

func IsBalanced(pRoot *TreeNode) bool {
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left)
		if left == -1 {
			return -1
		}
		right := depth(root.Right)
		if right == -1 {
			return -1
		}
		if abs(right-left) > 1 {
			return -1
		}
		return max(left, right) + 1
	}

	return depth(pRoot) > -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
