package tree

/*
二叉树中和为某一值的路径(三)
给定一个二叉树root和一个整数值 sum ，求该树有多少路径的的节点值之和等于 sum 。
1.该题路径定义不需要从根节点开始，也不需要在叶子节点结束，但是一定是从父亲节点往下到孩子节点
2.总节点数目为n
3.保证最后返回的路径个数在整形范围内(即路径个数小于231-1)
例如：
				  1
			2			3
		4       5   4       3
			-1
有三条路径符合要求：[2,4] [2,5,-1] [3,3]
*/

var pathsum3 = 0

func HasPathSum3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	dfsPathSum3(root, sum)
	HasPathSum3(root.Left, sum)
	HasPathSum3(root.Right, sum)
	return pathsum3
}

func dfsPathSum3(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		pathsum3++
	}
	dfsPathSum3(root.Left, sum)
	dfsPathSum3(root.Right, sum)
}
