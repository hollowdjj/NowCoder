package tree

import "nowcoder/utility"

/*
二叉树的直径
给定一颗二叉树，求二叉树的直径。
1.该题的直径定义为：树上任意两个节点路径长度的最大值
2.该题路径长度定义为：不需要从根节点开始，也不需要在叶子节点结束，也不需要必须从父节点到子节点，一个节点到底另外一个节点走的边的数目
3.这个路径可能穿过根节点，也可能不穿过
4.树为空时，返回 0
如，输入{1,2,3,#,#,4,5,9,#,#,6,#,7,#,8}，二叉树如下:
                1
		  2            3
                    4     5
                 9           6
                   7           8
那么路径长度最长为:7=>9=>4=>3=>5=>6=>8，长度为6
*/
func diameterOfBinaryTree(root *utility.TreeNode) int {
	//对一个二叉树节点而言，最长路径为其左子树的最大深度加上右子树的最大深度
	//因此，我们遍历每一个节点，即可得到最长路径
	res := 0
	var dfs func(node *utility.TreeNode) int
	dfs = func(node *utility.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return res
}
