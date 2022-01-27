package tree

/*
二叉搜索树转换为累加树
累加树是指每个节点的值是原来的节点值加上所有大于它的节点值之和。例如：
				5					18
			2		13   ---> 	20		13
*/
var preSum = 0

func BstToGst(root *TreeNode) *TreeNode {
	//遍历的顺序为：右-->根-->左，并且要记录右加根的和
	if root == nil {
		return root
	}
	BstToGst(root.Right)
	preSum += root.Val
	root.Val = preSum
	BstToGst(root.Left)
	return root
}
