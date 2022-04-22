package jz

/*
输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。（注意因为传入数据是链表，
所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）
*/

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	head1, head2 := pHead1, pHead2
	for pHead1 != pHead2 {
		if pHead1 == nil {
			pHead1 = head2
		} else {
			pHead1 = pHead1.Next
		}

		if pHead2 == nil {
			pHead2 = head1
		} else {
			pHead2 = pHead2.Next
		}
	}
	return pHead1
}
