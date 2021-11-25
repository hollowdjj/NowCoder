package list

import "nowcoder/utility"

/*
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。

要求：空间复杂度O(n),时间复杂度O(n)
进阶：空间复杂度O(1),时间复杂度O(n)
*/

func OddEvenList(head *utility.ListNode) *utility.ListNode {
	//节
	slice := make([]*utility.ListNode, 1)
	slice[0] = nil
	for head != nil {
		slice = append(slice, head)
		head = head.Next
	}
	//查询并分别生成奇偶链表
	dummyOdd, dummyEven := &utility.ListNode{-1, nil}, &utility.ListNode{-1, nil}
	tempOdd, tempEven := dummyOdd, dummyEven
	for i := 1; i < len(slice); i++ {
		if i%2 == 0 {
			//加入到偶数链表
			tempEven.Next = slice[i]
			tempEven = tempEven.Next
		} else {
			//加入到奇数链表
			tempOdd.Next = slice[i]
			tempOdd = tempOdd.Next
		}
	}
	//链接
	tempOdd.Next = dummyEven.Next
	tempEven.Next = nil
	return dummyOdd.Next
}
