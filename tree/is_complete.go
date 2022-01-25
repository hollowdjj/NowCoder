package tree

import "nowcoder/utility"

/*
判断二叉树是否是一颗完全二叉树
*/

func IsCompleteTree(root *TreeNode) bool {
	//层序遍历直到遇到空节点
	var queue utility.Queue
	queue.Push(root)
	for front := (*queue.Front()).(*TreeNode); front != nil; {
		queue.Push(front.Left)
		queue.Push(front.Right)
	}
	//如果是一颗完全二叉树的话，此时队列里应该全是空节点
	for !queue.Empty() {
		front := (*queue.Front()).(*TreeNode)
		if front != nil {
			return false
		}
	}
	return true
}
