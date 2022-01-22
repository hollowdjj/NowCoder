package tree

import (
	"nowcoder/utility"
)

/*
将升序数组转化为平衡二叉搜索树
给定一个升序排序的数组，将其转化为平衡二叉搜索树（BST）。
平衡二叉搜索树指树上每个节点 node 都满足左子树中所有节点的的值都小于 node 的值，
右子树中所有节点的值都大于 node 的值，并且左右子树的节点数量之差不大于1

数据范围：0≤n≤10000，数组中每个值满足∣val∣≤5000
进阶：空间复杂度O(n)，时间复杂度O(n)
*/

func SortedArrayToBST(num []int) *utility.TreeNode {
	//二分法加递归即可
	n := len(num)
	if n == 0 {
		return nil
	} else if n == 1 {
		node := &utility.TreeNode{Val: num[0]}
		return node
	}
	mid := n / 2
	node := &utility.TreeNode{Val: num[mid]}
	node.Left = SortedArrayToBST(num[:mid])
	node.Right = SortedArrayToBST(num[mid:])
	return node
}
