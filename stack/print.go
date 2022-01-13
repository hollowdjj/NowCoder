package stack

import "nowcoder/utility"

/*
给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）

数据范围：0≤n≤1500,树上每个节点的val满足 |val| <= 100
要求：空间复杂度：O(n)，时间复杂度：O(n)
例如：
给定的二叉树是{1,2,3,#,#,4,5}
				1
              2   3
                 4  5
*/

func Print(pRoot *utility.TreeNode) [][]int {
	//考察的是层序遍历。但这里要使用两个栈，一个保存奇数层节点，另一个保存偶数层节点。
	//当前为奇数层时，所有节点的左子节点先入偶数层栈；当前层为偶数层时，所有节点的右子节点先入奇数层栈。
	if pRoot == nil {
		return [][]int{}
	}
	var stackOdd utility.Stack  //保存奇数层节点
	var stackEven utility.Stack //保存偶数层节点
	level := 1
	var result [][]int
	stackOdd.Push(pRoot)
	for !stackOdd.Empty() || !stackEven.Empty() {
		//当前层为奇数层
		if level%2 == 1 {
			var res []int
			for !stackOdd.Empty() {
				temp := (*stackOdd.Pop()).(*utility.TreeNode)
				res = append(res, temp.Val)
				//奇数层所有节点的左子节点先入偶数层栈
				if temp.Left != nil {
					stackEven.Push(temp.Left)
				}
				if temp.Right != nil {
					stackEven.Push(temp.Right)
				}
			}
			if len(res) > 0 {
				result = append(result, res)
				level++
			}
		} else {
			//当前层为偶数层
			var res []int
			for !stackEven.Empty() {
				temp := (*stackEven.Pop()).(*utility.TreeNode)
				res = append(res, temp.Val)
				//偶数层所有子节点的右子节点先入奇数层栈
				if temp.Right != nil {
					stackOdd.Push(temp.Right)
				}
				if temp.Left != nil {
					stackOdd.Push(temp.Left)
				}
			}
			if len(res) > 0 {
				result = append(result, res)
				level++
			}
		}
	}
	return result
}
