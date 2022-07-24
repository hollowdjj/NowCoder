package leetcode

/*
给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。

输入: head = [-10,-3,0,5,9]
输出: [0,-3,9,-10,null,5]
解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	var work func(head, end *ListNode) *TreeNode
	work = func(head, end *ListNode) *TreeNode {
		if head == end {
			return nil
		}
		mid := findMid(head, end)
		root := &TreeNode{Val: mid.Val}
		root.Left = work(head, mid)
		root.Right = work(mid.Next, end)
		return root
	}
	return work(head, nil)
}

func findMid(head, end *ListNode) *ListNode {
	//找到[head,end]的中点
	slow, fast := head, head
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
