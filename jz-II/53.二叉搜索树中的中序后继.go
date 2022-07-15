package jz_II

/*
给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。

节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
*/

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//中序遍历
	hit, found := false, false
	var res *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || found {
			return
		}
		inorder(root.Left)
		if !found && hit {
			found = true
			res = root
			return
		}
		if root == p {
			hit = true
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}

func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	//利用二叉搜索树的性质：
	//1.如果p有右子树，那么中序后继就是p的右子树中值最小的节点
	//2.如果p没有右子树，那么中序后继是p所在的子树中大于它的最小值
	var res *TreeNode
	if node := p.Right; node != nil {
		res = node
		for node.Left != nil {
			node = node.Left
			res = node
		}
		return res
	}
	for root != nil {
		if root.Val > p.Val {
			res = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return res
}
