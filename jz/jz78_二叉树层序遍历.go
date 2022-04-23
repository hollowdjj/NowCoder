package jz

/*
二叉树层序遍历
*/

func LevelOrder(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	res := make([][]int, 0)
	q := []*TreeNode{pRoot}
	for len(q) > 0 {
		size := len(q)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			top := q[0]
			q = q[1:]
			temp[i] = top.Val
			if top.Left != nil {
				q = append(q, top.Left)
			}
			if top.Right != nil {
				q = append(q, top.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
