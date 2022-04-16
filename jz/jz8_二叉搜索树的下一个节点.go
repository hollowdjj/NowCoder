package jz

/*
给定一个二叉树其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的
next指针。例如：
					8
			6				10
		5		7		9		11
*/

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	//分两种情况：
	//1. 当该节点有右子树时，中序遍历的下一个节点就是这个右子树最左边的那个节点
	//2. 当该节点没有右子树时，又分为了两种情况：
	//2.1 当该节点为父节点的左子节点时，下一个节点就是父节点
	//2.2 当该节点为父节点的右子节点时，我们需要一直往上遍历，直到找到某一个节点，其是其父节点的左子节点。
	if pNode.Right != nil {
		pNode = pNode.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		return pNode
	}

	for pNode.Next != nil {
		parent := pNode.Next
		if pNode == parent.Left {
			return parent
		}
		pNode = pNode.Next
	}

	return nil
}
