package tree

import (
	"math"
	"nowcoder/utility"
)

/*
判断一棵二叉树是否为搜索二叉树和完全二叉树

给定一棵二叉树，已知其中的节点没有重复值，请判断该二叉树是否为搜索二叉树和完全二叉树。
输出描述：分别输出是否为搜索二叉树、完全二叉树。
数据范围：二叉树节点数满足 0≤n≤500000 ，二叉树上的值满足 0≤val≤100000
要求：空间复杂度O(n)，时间复杂度O(n)
注意：空子树我们认为同时符合搜索二叉树和完全二叉树。

这里总结一下各种各样的二叉树：
1. 满二叉树。即叶子节点只在最后一层有，其节点个数为2^height - 1
				  1
              2       3
           4     5  6    7
2. 完全二叉树。即叶子节点只在最后两层有，且在最后一层中，叶子节点从左到有依次排列，中间不能有空节点
                  1
              2       3
           4     5  6
下面这个就不是完全二叉树
                  1
              2       3
           4     5       7
3. 二叉搜索树。满足左子树中的所有节点值都小于根节点，而右子树中的所有值都大于根节点。也就是说，对二叉
搜索树进行中序遍历，得到的数组是一个升序数组[1,2,3,5,6,8,9]。
                   5
              2         8
           1     3  6		9
4. 平衡二叉树。即左右子树的高度差小于等于1的二叉树
                  1
              2       3
           4     5
*/

func Judge(root *TreeNode) []bool {
	res := make([]bool, 2)
	res[0] = isSearchTree(root)
	res[1] = isCompleteTree(root)
	return res
}

var pre = math.MinInt

func isSearchTree(root *TreeNode) bool {
	//中序遍历判断是否是二叉搜索树
	if root == nil {
		return true
	}
	//判断左子树
	if !isSearchTree(root.Left) {
		return false
	}
	if root.Val < pre {
		return false
	}
	//判断右子树
	pre = root.Val
	return isSearchTree(root.Right)
}

func isCompleteTree(root *TreeNode) bool {
	//层序遍历直到遇到空节点
	var q utility.Queue
	q.Push(root)
	for (*q.Front()).(*TreeNode) != nil {
		front := (*q.Pop()).(*TreeNode)
		q.Push(front.Left)
		q.Push(front.Right)
	}
	//如果是一个完全二叉树的话，此时队列里面应该全部是空节点
	for !q.Empty() {
		front := (*q.Pop()).(*TreeNode)
		if front != nil {
			return false
		}
	}
	return true
}
