package jz

import "math"

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回 true ,否则返回 false 。
假设输入的数组的任意两个数字都互不相同。
*/

func VerifySequenceOfBST(postorder []int) bool {
	if len(postorder) == 0 {
		return false
	}
	//后序遍历顺序是：左右根，倒序遍历后续遍历数组即：根右左
	//那么，将根节点和右节点的值入栈
	stack := make([]int, 0)
	root := math.MaxInt32
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && postorder[i] < stack[len(stack)-1] {
			//在栈中找到大于postorder[i]的最小值，该值就是根节点
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
