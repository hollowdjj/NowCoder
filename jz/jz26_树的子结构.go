package jz

/*
输入两棵二叉树A，B，判断B是不是A的子结构。（我们约定空树不是任意一个树的子结构）
假如给定A为{8,8,7,9,2,#,#,#,#,4,7}，B为{8,9,2}，2个树的结构如下，可以看出B是A的子结构

			8								8
	 8				7					9		2
 9       2
     4       7
*/

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if IsSame(pRoot1, pRoot2) {
		return true
	}

	return HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func IsSame(root1, root2 *TreeNode) bool {
	left, right := true, true
	//因为是以root2的结构进行的递归，所以一旦root1为空，或者两者值不相等，那么一定
	//没有相同的结构
	if root1 == nil || root1.Val != root2.Val {
		return false
	}
	if root2.Left != nil {
		left = IsSame(root1.Left, root2.Left)
	}
	if root2.Right != nil {
		right = IsSame(root1.Right, root2.Right)
	}
	return left && right
}
