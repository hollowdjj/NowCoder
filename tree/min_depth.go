package tree

import "nowcoder/utility"

/*
二叉树的最小深度
给定一颗节点数为N的二叉树，求其最小深度。最小深度是指树的根节点到最近叶子节点的最短路径上节点的数量。
（注：叶子节点是指没有子节点的节点。）
						1
					2	    3
 				4       5
上面二叉树最小深度为2
*/

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//层序遍历，遇到叶子节点就返回
	level := 0
	var queue utility.Queue
	queue.Push(root)
	for !queue.Empty() {
		size := queue.Size()
		for ; size > 0; size-- {
			front := (*queue.Pop()).(*TreeNode)
			if front.Right == nil && front.Left == nil {
				return level
			}
			if front.Left != nil {
				queue.Push(front.Left)
			}
			if front.Right != nil {
				queue.Push(front.Right)
			}
		}
		level++
	}
	return level
}
