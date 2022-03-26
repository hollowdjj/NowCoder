package tree

/*
给定一个二叉树的中序与后序遍历结果，请你根据两个序列构造符合这两个序列的二叉树。

例如输入[2,1,4,3,5],[2,4,5,3,1]时，
根据中序遍历的结果[2,1,4,3,5]和后序遍历的结果[2,4,5,3,1]可构造出二叉树{1,2,3,#,#,4,5}，如下图所示：
							1
						2		3
							4		5
*/

func BuildTree(inorder []int, postorder []int) *TreeNode {
	//定义一个map，key为数字，value为该数字在中序遍历数组中的索引
	dic := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}

	return work(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, dic)
}

func work(inorder []int, inLeft, inRight int, postorder []int, postLeft, postRight int, dic map[int]int) *TreeNode {
	if inLeft > inRight {
		return nil
	}

	//根节点为后续遍历数组的最后一个数字
	root := &TreeNode{Val: postorder[postRight]}

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

/*
给定一个二叉树中序与前序遍历结果，请你根据两个序列构造符合这两个序列的二叉树。
例如输入[2,1,4,3,5],[1,2,3,4,5]时，构造出二叉树：
							1
						2		3
							4		5
*/

func BuildTreeByInorderAndPrevOrder(inorder []int, preorder []int) *TreeNode {
	//使用一个哈希表保存中序遍历数组的结果与其对应的索引
	dic := make(map[int]int)
	for i, num := range inorder {
		dic[num] = i
	}

	return work2(inorder, 0, len(inorder)-1, preorder, 0, len(preorder)-1, dic)
}

func work2(inorder []int, inLeft, inRight int, preorder []int, preLeft, preRight int, dic map[int]int) *TreeNode {
	//递归终止条件
	if inLeft > inRight || preLeft > preRight {
		return nil
	}
	//根节点
	root := &TreeNode{Val: preorder[preLeft]}
	//前序遍历数组的第一个元素为根节点，查找出其在中序遍历数组中的索引
	index := dic[preorder[preLeft]]
	//左子树节点个数
	leftNum := index - inLeft
	//递归构建左子树和右子树
	root.Left = work2(inorder, inLeft, index-1, preorder, preLeft+1, leftNum+preLeft, dic)
	root.Right = work2(inorder, index+1, inRight, preorder, leftNum+preLeft+1, preRight, dic)
	return root
}

/*
给定一个二叉树前序和后序遍历的结果，请你根据两个序列构造符合这两个序列的二叉树。
例如输入[1,2,3,4,5]和[2,4,5,3,1]，构造出二叉树：
							1
						2		3
							4		5
*/
func BuildTreeByPrevOrderAndPostOrder(pre []int, post []int) *TreeNode {
	//记录post的中的元素与其索引
	dic := make(map[int]int)
	for i, num := range post {
		dic[num] = i
	}
	return work3(pre, 0, len(pre)-1, post, 0, len(post)-1, dic)
}

func work3(pre []int, preLeft, preRight int, post []int, postLeft, postRight int, dic map[int]int) *TreeNode {
	//递归终止条件
	if preLeft > preRight {
		return nil
	}

	//根节点
	root := &TreeNode{Val: pre[preLeft]}
	if preLeft == preRight {
		return root
	}

	//计算出左子树的元素个数
	index := dic[pre[preLeft+1]] //左子树根节点在后序遍历数组中的位置
	leftNum := index - postLeft + 1

	//递归构建左子树和右子树
	root.Left = work3(pre, preLeft+1, preLeft+leftNum, post, postLeft, index, dic)
	root.Right = work3(pre, preLeft+leftNum+1, preRight, post, index+1, postRight-1, dic)
	return root
}
