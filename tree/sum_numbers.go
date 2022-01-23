package tree

/*
二叉树根节点到叶子节点的所有路径和
给定一个二叉树的根节点root，该树的节点值都在数字0−9之间，每一条从根节点到叶子节点的路径都可以用一个数字表示。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n
例如根节点到叶子节点的一条路径是1→2→3,那么这条路径就用123来代替。
例如：
				1
			2       3
有两条路径1->2，用数字12代替；1->3，用数字13代替。所以答案为12+13=25
*/

func SumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			temp := num*10 + node.Val
			sum += temp
			return
		}
		num *= 10
		num += node.Val
		dfs(node.Left, num)
		dfs(node.Right, num)
	}
	dfs(root, 0)
	return sum
}
