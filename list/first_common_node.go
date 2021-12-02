package list

import "nowcoder/utility"

/*
输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。
要求：空间复杂度 O(1)，时间复杂度O(n)
*/

//FindFirstCommonNode 找到两个无环单向链表的公共节点
func FindFirstCommonNode(pHead1 *utility.ListNode, pHead2 *utility.ListNode) *utility.ListNode {
	/*
		这道题的解法比较巧妙，之前没见过的话，感觉很难想到。设链表1的头结点到公共结点的
		距离为a，公共结点到尾结点的距离为c;同样对链表2有距离b,c。那么显然有：
						a + c + b + c = b + c + a + c
		其中c部分的第一个结点就是公共结点。也就是说，将链表2接到链表1的后面，然后将链表1
		接到链表2的后面，然后链表1,2同时从头结点开始向后移动一格，就一定能找到公共结点
	*/
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
