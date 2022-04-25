package jz

import "testing"

func TestFindContinuousSequence(t *testing.T) {
	FindContinuousSequence(9)
}

func TestSerializeBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	DeserializeFromLevelOrder(SerializeByLevelOrder(root))
}
