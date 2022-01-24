package tree

/*
完全二叉树结点数
给定一棵完全二叉树的头节点head，返回这棵树的节点个数。
完全二叉树指：设二叉树的深度为h，则 [1,h-1] 层的节点数都满足 2^{i-1}个

数据范围：节点数量满足 0≤n≤100000，节点上每个值都满足 0≤val≤100000

对一颗完全二叉树而言，当左子树的高度大于右子树的高度时，右子树一定是一颗满二叉树。
而当左子树的高度等于右子树的高度时，左子树一定是一颗满二叉树。
                  1
              2       3
           4     5  6
*/

func NodeNum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)
	if left == right {
		return 1 + 1<<left - 1 + NodeNum(root.Right)
	}
	return 1 + 1<<right - 1 + NodeNum(root.Left)
}

func getHeight(root *TreeNode) int {
	//完全二叉树的深度只取决于左子树
	res := 0
	for root != nil {
		res++
		root = root.Left
	}
	return res
}
