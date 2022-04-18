package jz

/*
输入一颗二叉树的根节点root和一个整数expectNumber，找出二叉树中结点值的和为expectNumber的所有路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

				10
			5		12
		4		7
expectNumber为22的合法路径有[10,5,7]和[10,12]
*/

func FindPath(root *TreeNode, expectNumber int) [][]int {
	res := make([][]int, 0)
	var dfs func(root *TreeNode, curr []int, rest int)
	dfs = func(root *TreeNode, curr []int, rest int) {
		if root == nil {
			return
		}
		//因为是从头节点到叶子节点，所以要先添加值，然后判断条件为rest == root.Val
		curr = append(curr, root.Val)
		if root.Left == nil && root.Right == nil {
			if rest == root.Val {
				temp := make([]int, len(curr))
				copy(temp, curr)
				res = append(res, temp)
			}
			return
		}
		dfs(root.Left, curr, rest-root.Val)
		dfs(root.Right, curr, rest-root.Val)
	}
	dfs(root, []int{}, expectNumber)
	return res
}
