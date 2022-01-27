package tree

import "math"

/*
二叉树中的最大路径和
二叉树里面的路径被定义为:从该树的任意节点出发，经过父=>子或者子=>父的连接，达到任意节点的序列。
注意:
1.同一个节点在一条二叉树路径里中最多出现一次
2.一条路径至少包含一个节点，且不一定经过根节点
给定一个二叉树的根节点root，请你计算它的最大路径和
例如：
					-20
				8		  20
					 15			6
最大路径和为15+20+6=41
*/
var preMax = math.MinInt

func MaxPathSum(root *TreeNode) int {
	//这道题和“二叉树的直径”那道题有点相似。“二叉树的直径”求得是左右子树的最大
	//深度的和。而本题其实就是要求左右子树节点相加的最大和。也就是说，这道题的
	//本质就是让你求“以二叉树中任何一个节点为根节点的树的左右子树最大路径和的最大
	//值”
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		//计算左右子树的最大路径和。出现负数时为0
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		//对于根节点为root的树而言，其最大路径和为left+right+root.Val
		preMax = max(preMax, left+right+root.Val)
		/*
				这里的返回值需要特别注意。返回值只能时root的值加上左右子树最大路径
				和中的较大值。因为路径是不能出现分叉的，例如路径不能是1 3 (6,7)
									1
								2		3
			                 4	   5 6	    7
		*/
		return max(0, root.Val+max(left, right))
	}
	dfs(root)
	return preMax
}
