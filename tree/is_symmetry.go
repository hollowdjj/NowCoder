package tree

import "nowcoder/utility"

/*
判断二叉树是不是一颗对称二叉树

要求：时间复杂度O(n)，空间复杂度O(n)
*/

func IsSymmetry(root *utility.TreeNode) bool {
	/*
		如下图所示，当一颗二叉树是对称二叉树时。对一个节点的两个子节点left和right有如下要求：
		2. left.Left.Val == right.Right.Val && left.Right.Val == right.Left.Val
							1
						2       2
					 3     4 4     3
		需要注意的是，此题不能用中序遍历求解，因为：
			                1
			            3       2
			         2        3
		此时，中序遍历的结果为[2,3,1,3,2]也是一个对称数组。
	*/
	return check(root, root)
}

func check(a, b *utility.TreeNode) bool {
	//递归终止条件
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && check(a.Left, b.Right) && check(a.Right, b.Left)
}
