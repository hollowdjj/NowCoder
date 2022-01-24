package tree

/*
二叉搜索树的第k个节点
给定一棵结点数为n 二叉搜索树，请找出其中的第 k 小的TreeNode结点值。
1.返回第k小的节点值即可
2.不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1
3.保证n个节点的值不一样

数据范围： 0≤n≤1000，0≤k≤1000，树上每个结点的值满足≤val≤1000
进阶：空间复杂度O(n)，时间复杂度O(n)
*/

var kthCount = 0

func KthNumber(root *TreeNode, k int) int {
	//显然是使用中序遍历
	if root == nil {
		return -1
	}
	//判断左子树
	res1 := KthNumber(root.Left, k)
	if res1 != -1 {
		return res1
	}
	kthCount++
	if kthCount == k {
		return root.Val
	}
	//判断右子树
	res2 := KthNumber(root.Left, k)
	return max(res1, res2)
}
