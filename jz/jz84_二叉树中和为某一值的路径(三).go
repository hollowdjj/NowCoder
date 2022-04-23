package jz

/*
给定一个二叉树root和一个整数值 sum ，求该树有多少路径的的节点值之和等于 sum 。
1.该题路径定义不需要从根节点开始，也不需要在叶子节点结束，但是一定是从父亲节点往下到孩子节点
2.总节点数目为n
3.保证最后返回的路径个数在整形范围内(即路径个数小于231-1)
*/

var res84 = 0

func FindPath3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	dfs84(root, sum)
	FindPath(root.Left, sum)
	FindPath(root.Right, sum)
	return res84
}

func dfs84(root *TreeNode, rest int) {
	if root == nil {
		return
	}
	rest -= root.Val
	if rest == 0 {
		res84++
	}
	dfs84(root.Left, rest)
	dfs84(root.Right, rest)
}
