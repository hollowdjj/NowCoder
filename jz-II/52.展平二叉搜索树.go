package jz_II

func IncreasingBST(root *TreeNode) *TreeNode {
	res := &TreeNode{Val: -1}
	head := res
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		head.Right = root
		root.Left = nil
		head = head.Right
		inorder(root.Right)
	}
	inorder(root)
	return res.Right
}
