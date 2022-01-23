package tree

import "nowcoder/utility"

/*
二叉树中和为某一值的路径(二)
输入一颗二叉树的根节点root和一个整数expectNumber，找出二叉树中结点值的和为expectNumber的所有路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n
*/

func HasPathSum2(root *utility.TreeNode, expectNumber int) [][]int {
	var res [][]int
	var dfs func(node *utility.TreeNode, rest int, path []int)
	dfs = func(node *utility.TreeNode, rest int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && rest == node.Val {
			/*
				这里需要注意的是Go中切片的一个坑。在Go中，切片是一个引用类型，
				也就是说，在作为函数的参数传递时，实际上传递的是一个指针，即地址。
						                          5
										      4       8
					                       11       13  4
					                      7  2         5  1
				当我们得到了答案[5,8,4,5]时，此时底层数组为[5,8,4,5]。然后，在回溯并
				开始递归4的右子树时,path引用的是底层数组的前三个元素，即[5,8,4]。此时，
				再append 1进去，由于底层数组还有空间，从而底层数组变成了[5,8,4,1]。
				进而导致结果错误。
			*/
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		dfs(node.Left, rest-node.Val, path)
		dfs(node.Right, rest-node.Val, path)
	}
	dfs(root, expectNumber, []int{})
	return res
}
