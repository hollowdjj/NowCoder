package tree

import "nowcoder/utility"

/*
二叉树的层序遍历
*/

func LevelOrder(root *utility.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue utility.Queue
	queue.Push(root)
	for !queue.Empty() {
		//层序遍历中的关键点在于判断哪些节点处于同一层。每个while循环就是一层，因此，
		//在while循环的伊始，记录一下当前层的元素个数即可。
		size := queue.Size()
		level := make([]int, size)
		for i := 0; i < size; i++ {
			front := (*queue.Pop()).(*utility.TreeNode)
			level[i] = front.Val
			if front.Left != nil {
				queue.Push(front.Left)
			}
			if front.Right != nil {
				queue.Push(front.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
