package tree

/*
二叉树的中序遍历
*/

func Inorder(root *TreeNode) []int {
	var res []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

//非递归实现中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		//一直走左子树直到叶子节点
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		//          1
		//      2       3
		//   4     5
		//例如，得到叶子节点4，打印值，然后处理它的右子节点
		//右子节点为空，所以不会走上面那个while循环。然后打
		//印2，并处理2的右子树。
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, top.Val) //根节点
			root = top.Right
		}
	}
	return res
}
