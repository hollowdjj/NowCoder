package codetop

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//注意，遇到像这种需要保存前置节点的情况，最好的做法就是声明一个dummy节点。
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	curr := prev.Next
	for i := left; i < right; i++ {
		//头插法的应用，一定要理解并牢记。将curr.Next和prev.Next交换
		//1 2 3 4 5 --> 1 3 2 4 5  --> 1 4 3 2 5
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}
