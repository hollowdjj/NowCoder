package tree

import "nowcoder/utility"

/*
判断二叉树是不是平衡二叉树
在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的
高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。

约定空树也是一颗平衡二叉树
要求：空间复杂度O(1)，时间复杂度O(n)
*/

func IsBalanced(root *utility.TreeNode) bool {
	//思路就是对二叉树中的每一个节点计算以其为根节点的二叉树的左右节点
	//深度差是否大于1。
	if root == nil {
		return true
	}
	ld := MaxDepth(root.Left)
	rd := MaxDepth(root.Right)
	if abs(ld-rd) > 1 {
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

func IsBalancedAdvanced(root *utility.TreeNode) bool {
	//我们注意到，IsBalanced函数是遍历了二叉树中的所有节点。然而，事实上
	//只要有一个节点不满足平衡二叉树的条件，那么这颗二叉树就不是一颗平衡
	//二叉树。因此，我们修改求解二叉树最大深度的函数，在节点的左右子树高度
	//差大于1时，返回-1，从而实现了剪枝。
	return depth(root) > -1
}

func depth(node *utility.TreeNode) int {
	if node == nil {
		return 0
	}
	ld := depth(node.Left)
	if ld == -1 {
		return -1
	}
	rd := depth(node.Right)
	if rd == -1 {
		return -1
	}
	if abs(ld-rd) > 1 {
		return -1
	}
	return max(ld, rd) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
