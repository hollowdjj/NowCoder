package jz_II

/*
给定一个链表的 头节点 head ，请判断其是否为回文链表。
*/

func isPalindrome2(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	//快慢指针找到链表中点
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//断开并翻转后半部分链表
	prev.Next = nil
	head1 := reverseList(slow)
	//判断
	for head != nil && head1 != nil {
		if head.Val != head1.Val {
			return false
		}
		head = head.Next
		head1 = head1.Next
	}

	return true
}
