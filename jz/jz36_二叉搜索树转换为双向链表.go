package jz

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示:
          10
	6		 	14       ==>    4≒6≒8≒10≒12≒14≒16
4       8   12      16

*/
var prev36 *TreeNode

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}

	inorder(pRootOfTree)
	//一直往左走到头
	for pRootOfTree.Left != nil {
		pRootOfTree = pRootOfTree.Left
	}
	return pRootOfTree
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	root.Left = prev36
	if prev36 != nil {
		prev36.Right = root
	}
	prev36 = root
	inorder(root.Right)
}
