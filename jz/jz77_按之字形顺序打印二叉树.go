package jz

/*
给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）
*/

func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	//这道题使用的是“双栈”，不是一个栈和一个队列！！！！！！！！
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)
	stack1 = append(stack1, pRoot)
	level := 1
	res := make([][]int, 0)
	for len(stack1) > 0 || len(stack2) > 0 {
		if level%2 == 1 {
			//奇数层使用栈1，并且左子节点先入栈2
			temp := make([]int, 0)
			for len(stack1) > 0 {
				top := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				temp = append(temp, top.Val)
				if top.Left != nil {
					stack2 = append(stack2, top.Left)
				}
				if top.Right != nil {
					stack2 = append(stack2, top.Right)
				}
			}
			res = append(res, temp)
		} else {
			//偶数层使用栈2，并且右子节点先入栈1
			temp := make([]int, 0)
			for len(stack2) > 0 {
				top := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				temp = append(temp, top.Val)
				if top.Right != nil {
					stack1 = append(stack1, top.Right)
				}
				if top.Left != nil {
					stack1 = append(stack1, top.Left)
				}
			}
			res = append(res, temp)
		}
		level++
	}
	return res
}
