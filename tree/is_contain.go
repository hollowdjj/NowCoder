package tree

import "nowcoder/utility"

/*
给定彼此独立的两棵二叉树，树上的节点值两两不同，判断 t1 树是否有与 t2 树完全相同的子树。
子树指一棵树的某个节点的全部后继节点

数据范围：树的节点数满足 0<n≤500000，树上每个节点的值一定在32位整型范围内
进阶：空间复杂度: O(1)，时间复杂度O(n)
*/

func IsContain(root1 *utility.TreeNode, root2 *utility.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return IsContain(root1.Left, root2.Left) && IsContain(root1.Right, root2.Right)
	}
	return IsContain(root1.Left, root2) || IsContain(root1.Right, root2)
}
