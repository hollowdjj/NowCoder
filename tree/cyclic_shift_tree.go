package tree

import "nowcoder/utility"

/*
循环右移二叉树
现有一棵nn个节点构成的二叉树，请你将每一层的节点向右循环位移kk位。某层向右位移一位(即k=1k=1)的含义为：
1.若当前节点为左孩子节点，会变成当前节点的双亲节点的右孩子节点。
2.若当前节点为右儿子，会变成当前节点的双亲节点的右边相邻兄弟节点的左孩子节点。
(如果当前节点的双亲节点已经是最右边的节点了，则会变成双亲节点同级的最左边的节点的左孩子节点)
3.该层的每一个节点同时进行一次位移。
4.是从最下面的层开始位移，位移完每一层之后，再向上，直到根节点，位移完毕。
例如：
					1								1
				2		3			---->		3		2
					4		5					  4   5
*/

func CyclicShiftTree(root *TreeNode, k int) *TreeNode {
	//思路很简单，就是保存每一层包括空节点在内的所有节点。然后，对每一层进行
	//右移，然后让上一层重新挂载节点
	level := make([][]*TreeNode, 0)
	var q utility.Queue
	q.Push(root)
	for !q.Empty() {
		size := q.Size()
		temp := make([]*TreeNode, size)
		for i := 0; i < size; i++ {
			front := (*q.Pop()).(*TreeNode)
			temp[i] = front
			if front != nil {
				q.Push(front.Left)
				q.Push(front.Right)
			}
		}
		level = append(level, temp)
	}
	level = level[:len(level)-1]
	n := len(level)
	for i := n - 1; i > 0; i-- {
		size := len(level[i])
		m := k % size //得到等效的移动次数m(0≤m﹤size)
		//循环右移m次，就是把最后m个元素放到数组头部
		f := make([]*TreeNode, size-m)
		copy(f, level[i][:size-m])
		b := make([]*TreeNode, m)
		copy(b, level[i][size-m:])
		index := 0
		for j, _ := range b {
			level[i][index] = b[j]
			index++
		}
		for j, _ := range f {
			level[i][index] = f[j]
			index++
		}
		//重新挂载第i-1层的子节点
		index = 0
		for j, _ := range level[i-1] {
			if level[i-1][j] != nil {
				level[i-1][j].Left = level[i][index]
				index++
				level[i-1][j].Right = level[i][index]
				index++
			}
		}
	}
	return root
}
