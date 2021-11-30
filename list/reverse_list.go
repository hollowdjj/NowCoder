package list

import "nowcoder/utility"

//翻转链表

func ReverseList(pHead *utility.ListNode) *utility.ListNode {
	var prev *utility.ListNode
	for pHead != nil {
		next := pHead.Next //先把后一个结点保存了
		pHead.Next = prev  //修改当前结点后后续结点为当前结点的前一个结点
		prev = pHead       //update
		pHead = next       //update
	}
	return prev
}
