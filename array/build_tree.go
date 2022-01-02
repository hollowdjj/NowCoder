package array

import "nowcoder/utility"

/*
给定一个二叉树的中序与后序遍历结果，请你根据两个序列构造符合这两个序列的二叉树。

例如输入[2,1,4,3,5],[2,4,5,3,1]时，
根据中序遍历的结果[2,1,4,3,5]和后序遍历的结果[2,4,5,3,1]可构造出二叉树{1,2,3,#,#,4,5}，如下图所示：
*/

func BuildTree(inorder []int, postorder []int) *utility.TreeNode {
	//定义一个map，key为数字，value为该数字在中序遍历数组中的索引
	dic := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}

	return work(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, dic)
}

func work(inorder []int, inLeft, inRight int, postorder []int, postLeft, postRight int, dic map[int]int) *utility.TreeNode {
	if inLeft > inRight {
		return nil
	}

	//根节点为后续遍历数组的最后一个数字
	root := &utility.TreeNode{Val: postorder[postRight]}

	//查询dic，找到root在中序遍历数组中的索引
	index := dic[root.Val]

	//根据中序遍历数组计算根节点左子树以及右子树的元素数量
	leftNum := index - inLeft
	//rightNum := inRight - index

	//递归构建左子树和右子树
	root.Left = work(inorder, inLeft, index-1, postorder, postLeft, postLeft+leftNum-1, dic)
	root.Right = work(inorder, index+1, inRight, postorder, postLeft+leftNum, postRight-1, dic)

	return root
}
