package tree

import "nowcoder/utility"

/*
修剪叶子
有一棵有n个节点的二叉树，其根节点为root。修剪规则如下:
1.修剪掉当前二叉树的叶子节点，但是不能直接删除叶子节点
2.只能修剪叶子节点的父节点，修剪了父节点之后，叶子节点也会对应删掉
3.如果想在留下尽可能多的节点前提下，修剪掉所有的叶子节点。请你返回修剪后的二叉树。
*/

func PruneLeaves(root *utility.TreeNode) *utility.TreeNode {
	if root == nil {
		return nil
	}
	//如果左孩子是叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return nil
	}
	//如果右孩子是叶子节点
	if root.Right != nil && root.Right.Left == nil && root.Right.Right == nil {
		return nil
	}
	root.Left = PruneLeaves(root.Left)
	root.Right = PruneLeaves(root.Right)
	return root
}
