package jz_II

/*
给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素insertVal ，使这个列表
仍然是循环升序的。给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(node *Node, x int) *Node {
	newNode := &Node{Val: x}
	//若链表为空，创建循环链表
	if node == nil {
		newNode.Next = newNode
		return newNode
	}
	//若链表只有一个节点，插入前后都行
	if node.Next == node {
		node.Next = newNode
		newNode.Next = node
		return node
	}
	//找到一个节点，其值小于等于x，但后一个节点大于等于x，插入节点
	curr, maxNode := node, node
	for {
		if curr.Val <= x && curr.Next.Val >= x {
			newNode.Next = curr.Next
			curr.Next = newNode
			return node
		}
		curr = curr.Next
		if curr.Val >= maxNode.Val {
			maxNode = curr
		}
		if curr == node {
			break
		}
	}
	//没有找到插入位置，说明x比链表中的最大值大或者比最小值小。
	//这两种情况，都只需要将x插入到链表最大值节点的后面
	newNode.Next = maxNode.Next
	maxNode.Next = newNode
	return node
}
