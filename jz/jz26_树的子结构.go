package jz

/*
输入两棵二叉树A，B，判断B是不是A的子结构。（我们约定空树不是任意一个树的子结构）
假如给定A为{8,8,7,9,2,#,#,#,#,4,7}，B为{8,9,2}，2个树的结构如下，可以看出B是A的子结构

			8								8
	 8				7					9		2
 9       2
     4       7
*/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

//判断root2是不是root1的子结构。root1和root2不能同时为nil
func isSub(root1, root2 *TreeNode) bool {
	//说明root2树已经匹配完成，满足子结构
	if root2 == nil {
		return true
	}
	//root2不为空，root1为空或者两者值不相等，说明不是子结构
	if root1 == nil || root1.Val != root2.Val {
		return false
	}
	return isSub(root1.Left, root2.Left) && isSub(root1.Right, root2.Right)
}
