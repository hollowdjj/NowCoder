package jz

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。返回删除后的链表的头节点。

1.此题对比原题有改动
2.题目保证链表中节点的值互不相同
3.该题只会输出返回的链表和结果做对比，所以若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
*/

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	prev, nHead := dummy, head
	for nHead != nil {
		if nHead.Val == val {
			prev.Next = nHead.Next
			break
		}
		prev = nHead
		nHead = nHead.Next
	}
	return dummy.Next
}
