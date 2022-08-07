package jz_II

/*
给定一个链表数组，每个链表都已经按升序排列。

请将所有链表合并到一个升序链表中，返回合并后的链表。

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
*/

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}

	return mergeTwo(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func mergeTwo(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			head.Next = head1
			head1 = head1.Next
		} else {
			head.Next = head2
			head2 = head2.Next
		}
		head = head.Next
	}
	if head1 != nil {
		head.Next = head1
	}
	if head2 != nil {
		head.Next = head2
	}
	return dummy.Next
}
