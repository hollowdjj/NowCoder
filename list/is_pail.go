package list

import "nowcoder/utility"

/*
给定一个链表，判断该链表是否是一个回文结构，即该链表的正逆序是否完全一致
*/

//IsPail 使用额外数组判断链表是否是一个回文结构。
func IsPail(head *utility.ListNode) bool {
	//将链表的值存放到一个切片中
	var s []int
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	//判断是否是回文结构
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//IsPailAdvanced 不适用额外数组判断链表是否是一个回文结构
func IsPailAdvanced(head *utility.ListNode) bool {
	/*
		首先找到链表的中间节点。找到链表中间节点的方法比较巧妙，利用的是快慢指针。慢指针一次走
		一步，快指针一次走两步。当快指针走到链表的最后一个节点时，慢指针刚好指向中间节点。当链
		表的节点数为偶数时，有两个中间节点，此时慢指针指向的是右边那个
	*/
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow
	//将链表后半部分反转。
	var prev *utility.ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	/*
		比较链表[head,prev]和链表[prev,mid]
	*/
	for head != mid {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}

	return true
}
