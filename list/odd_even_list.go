package list

import "nowcoder/utility"

/*
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。
例如：输入[1, 2, 3, 4, 5, 6]，返回[1, 3, 5,2, 4, 6]
要求：空间复杂度O(n),时间复杂度O(n)
进阶：空间复杂度O(1),时间复杂度O(n)
*/

//OddEvenList 将链表中奇数位节点和偶数位节点分别放在一起，重排后输出
func OddEvenList(head *utility.ListNode) *utility.ListNode {
	if head == nil {
		return head
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
