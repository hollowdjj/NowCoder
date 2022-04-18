package jz

/*
二叉树的层序遍历
*/

func PrintFromTopToBottom(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	res := make([]int, 0)
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		res = append(res, top.Val)
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
	}
	return res
}
