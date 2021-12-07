package array

import "nowcoder/utility"

/*
重建二叉树
给定节点数为 n 二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。
 							 1
                         /		 \
                        2         3
                       /         / \
                      4         5   6
                       \            /
                        7           8
要求：空间复杂度O(n)，时间复杂度O(n)
其他信息:
1.vin.length == pre.length
2.pre 和 vin 均无重复元素
3.vin出现的元素均出现在pre里
*/

func ReconstructBinaryTree(pre []int, vin []int) *utility.TreeNode {
	//key为二叉树节点值，value为其在中序遍历结果数组vin中的索引
	dic := make(map[int]int)
	for i, v := range vin {
		dic[v] = i
	}
	return construct(0, len(pre)-1, pre, 0, len(vin)-1, vin, dic)
}

func construct(preLeft, preRight int, preOrigin []int, vinLeft, vinRight int, vinOrigin []int, dic map[int]int) *utility.TreeNode {
	//递归终止条件
	if preLeft > preRight || vinLeft > vinRight {
		return nil
	}

	//根节点为前序遍历数组的首元素
	root := &utility.TreeNode{Val: preOrigin[preLeft]}
	//查询得到root在中序遍历数组中的索引index，从而左子树中元素的数量为index个
	index := dic[root.Val]
	//通过中序遍历数组vin计算得到根节点root的左子树的元素数量
	leftNum := index - vinLeft
	//递归计算左子节点
	root.Left = construct(preLeft+1, preLeft+leftNum, preOrigin, vinLeft, index-1, vinOrigin, dic)
	//递归计算右子节点
	root.Right = construct(preLeft+leftNum+1, preRight, preOrigin, index+1, vinRight, vinOrigin, dic)

	return root
}
