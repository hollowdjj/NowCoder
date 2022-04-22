package jz

/*
给定一棵结点数为n 二叉搜索树，请找出其中的第 k 小的TreeNode结点值。
1.返回第k小的节点值即可
2.不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1
3.保证n个节点的值不一样
*/

func KthNode(root *TreeNode, k int) int {
	res := -1
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || k < 0 {
			return
		}
		inorder(root.Left)
		k--
		if k == 0 {
			res = root.Val
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}
