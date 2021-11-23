package list

import "nowcoder/utility"

/*
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。给定两个这种链表，请
生成代表两个整数相加值的结果链表。
要求：空间复杂度 O(n)，时间复杂度 O(n)
例如：链表1为 9->3->7，链表2为 6->3，最后生成新的结果链表为 1->0->0->0。
*/

//AddInList 使用辅助栈。空间复杂度和时间复杂度均为O(n1+n2)
func AddInList(head1, head2 *utility.ListNode) *utility.ListNode {
	//使用两个辅助栈
	stack1, stack2 := utility.StackInt{}, utility.StackInt{}
	for head1 != nil {
		stack1.Push(head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		stack2.Push(head2.Val)
		head2 = head2.Next
	}
	//新链表的伪头结点
	head := &utility.ListNode{
		Val:  0,
		Next: nil,
	}
	//进位
	carry := 0
	//相加
	for !stack1.Empty() || !stack2.Empty() {
		sum := carry
		if !stack1.Empty() {
			sum += stack1.Pop()
		}
		if !stack2.Empty() {
			sum += stack2.Pop()
		}
		newNode := &utility.ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		//最新得到的节点插入到伪头结点head的后面
		low := head.Next
		head.Next = newNode
		newNode.Next = low
		carry = sum / 10
	}
	//最后一位可能发生进位
	if carry > 0 {
		low := head.Next
		newNode := &utility.ListNode{
			Val:  carry,
			Next: nil,
		}
		head.Next = newNode
		newNode.Next = low
	}

	return head.Next
}

//AddInListAdvanced 翻转链表后相加。空间复杂度为O(1)，时间复杂度为O(n)
func AddInListAdvanced(head1, head2 *utility.ListNode) *utility.ListNode {
	//翻转两链表
	head1, head2 = ReverseList(head1), ReverseList(head2)
	//进位
	carry := 0
	//新链表的伪头结点
	head := &utility.ListNode{
		Val:  -1,
		Next: nil,
	}
	//相加
	for head1 != nil || head2 != nil {
		sum := carry
		if head1 != nil {
			sum += head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			sum += head2.Val
			head2 = head2.Next
		}
		val := sum % 10
		carry = sum / 10
		//头插
		newHead := &utility.ListNode{
			Val:  val,
			Next: nil,
		}
		low := head.Next
		head.Next = newHead
		newHead.Next = low
	}
	//最后一位可能进位
	if carry > 0 {
		//头插
		low := head.Next
		newHead := &utility.ListNode{
			Val:  carry,
			Next: nil,
		}
		head.Next = newHead
		newHead.Next = low
	}

	return head.Next
}
