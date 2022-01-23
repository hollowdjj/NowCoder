package tree

/*
在二叉树中找到两个节点的最近公共祖先
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

数据范围：1≤n≤1000，树上每个节点的val满足 0<val≤100
要求：时间复杂度O(n)

注：本题保证二叉树中每个节点的val值均不相同。
*/
func LowestCommonAncestorByDfs(root *TreeNode, o1 int, o2 int) int {
	//使用DFS遍历所有节点，保存所有节点的父节点。一个更好的办法是使用层序遍历
	//一旦找到o1和o2对应的节点就停止。
	parents := make(map[int]*TreeNode)
	var dfs func(node *TreeNode, parent *TreeNode)
	dfs = func(node *TreeNode, parent *TreeNode) {
		if node == nil {
			return
		}
		parents[node.Val] = parent
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	//得到由o1到root以及o2到root的路径
	path1, path2 := []int{o1}, []int{o2}
	p1, p2 := parents[o1], parents[o2]
	for p1 != nil {
		path1 = append(path1, p1.Val)
		p1 = parents[p1.Val]
	}
	for p2 != nil {
		path2 = append(path2, p2.Val)
		p2 = parents[p2.Val]
	}
	//从后往前遍历path1 path2找到最后一个相同的数字
	n1, n2 := len(path1), len(path2)
	m := min(n1, n2)
	res := 0
	for i := 0; i < m; i++ {
		if path1[n1-1-i] == path2[n2-1-i] {
			res = path1[n1-1-i]
		}
	}
	return res
}

func LowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	//使用递归求解，平均时间复杂度要更优一些，但是不太好理解。需要结合图例理解。
	//						 3
	//					5		   1
	//				6       2   0      8
	//                    7   4
	var helper func(node *TreeNode, o1, o2 int) *TreeNode
	helper = func(node *TreeNode, o1, o2 int) *TreeNode {
		if node == nil || node.Val == o1 || node.Val == o2 {
			return node
		}
		//在node的左子树中找到o1或o2
		left := helper(node.Left, o1, o2)
		//在node的右子树中找到o1或o2
		right := helper(node.Right, o1, o2)
		//如果左子树中没有找到，说明o1和o2都在node的右子树上，那么最近的那个公共
		//祖先就是右子树的查询结果。相反，最近公共祖先为左子树的查询结果。
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		return node
	}
	return helper(root, o1, o2).Val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
