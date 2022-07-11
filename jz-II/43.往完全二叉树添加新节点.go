package jz_II

/*
设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为v。
使树保持完全二叉树的状态，并返回插入的新节点的父节点.
*/

type CBTInserter struct {
	root *TreeNode
}

func Constructor43(root *TreeNode) CBTInserter {
	return CBTInserter{root}
}

func (this *CBTInserter) Insert(v int) int {
	//层序遍历
	root := this.root
	q := []*TreeNode{root}
	res := 0
loop:
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:]
			if front.Left == nil {
				front.Left = &TreeNode{Val: v}
				res = front.Val
				break loop
			} else {
				q = append(q, front.Left)
			}
			if front.Right == nil {
				front.Right = &TreeNode{Val: v}
				res = front.Val
				break loop
			} else {
				q = append(q, front.Right)
			}
		}
	}
	return res
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
