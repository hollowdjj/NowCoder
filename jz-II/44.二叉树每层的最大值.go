package jz_II

import "math"

/*
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		n := len(q)
		maxVal := math.MinInt32
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:]
			if front.Val > maxVal {
				maxVal = front.Val
			}
			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}
		res = append(res, maxVal)
	}
	return res
}
