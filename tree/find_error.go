package tree

import "nowcoder/utility"

/*
找到搜索二叉树中两个错误的节点
一棵二叉树原本是搜索二叉树，但是其中有两个节点调换了位置，使得这棵二叉树不再是搜索二叉树，
请按升序输出这两个错误节点的值。(每个节点的值各不相同)
搜索二叉树：满足每个节点的左子节点小于当前节点，右子节点大于当前节点。
*/

func FindError(root *utility.TreeNode) []int {
	//中序遍历
	inorderRes := make([]int, 0)
	var inorder func(root *utility.TreeNode)
	inorder = func(root *utility.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		inorderRes = append(inorderRes, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	/*
		一个错误的二叉搜索树如下所示：
		           4
		       2       5
		   3      1
		其中序遍历的结果为：[3,2,1,4,5]。为了找到交换的那两个节点。首先逆序遍历该数组，找到
		交换的两点中的较小值(这里是1)。然后逆序遍历，找到较大值(这里是3)
	*/
	//首先逆序遍历中序数组得到较小值
	n := len(inorderRes)
	res := make([]int, 2)
	for i := n - 1; i > 0; i-- {
		if inorderRes[i] < inorderRes[i-1] {
			res[0] = inorderRes[i]
			break
		}
	}
	//然后顺序遍历中序数组得到较大值
	for i := 0; i < n-1; i++ {
		if inorderRes[i] > inorderRes[i+1] {
			res[1] = inorderRes[i]
			break
		}
	}
	return res
}
