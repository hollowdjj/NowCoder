package tree

/*
二叉搜索数转换为双向链表

输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示
			10
		6		14        ---->    4≒6≒8≒10≒12≒14≒16
	  4	  8   12   16
*/

var prevNode *TreeNode

func BstToList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	inorder(root)
	//一直往左遍历找到链表的头结点
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func inorder(curr *TreeNode) {
	if curr == nil {
		return
	}
	inorder(curr.Left)
	curr.Left = prevNode
	if prevNode != nil {
		prevNode.Right = curr
	}
	prevNode = curr
	inorder(curr.Right)
}
