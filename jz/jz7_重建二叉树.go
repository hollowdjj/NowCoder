package jz

/*
给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}
*/

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	dic := make(map[int]int)
	for i, v := range vin {
		dic[v] = i
	}
	return work(pre, 0, len(pre)-1, vin, 0, len(vin)-1, dic)
}

func work(pre []int, preLeft, preRight int, vin []int, vinLeft, vinRight int, dic map[int]int) *TreeNode {
	if preLeft > preRight || vinLeft > vinRight {
		return nil
	}
	root := &TreeNode{Val: pre[preLeft]}
	index := dic[pre[preLeft]]
	leftNum := index - vinLeft
	root.Left = work(pre, preLeft+1, leftNum+preLeft, vin, vinLeft, index-1, dic)
	root.Right = work(pre, leftNum+preLeft+1, preRight, vin, index+1, vinRight, dic)
	return root
}
