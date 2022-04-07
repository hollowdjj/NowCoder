package list

import "nowcoder/utility"

/*
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表。如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样。
你不能更改节点中的值，只能更改节点本身。
要求：空间复杂度 O(1)，时间复杂度 O(n)
例如：
给定的链表是1→2→3→4→5
对于k=2,你应该返回2→1→4→3→5
对于k=3,你应该返回3→2→1→4→5
*/

func ReverseKGroup(head *utility.ListNode, k int) *utility.ListNode {
	if head == nil || k == 1 {
		return head
	}
	left, right := head, head
	var tail *utility.ListNode
	count := 1
	for right.Next != nil {
		right = right.Next
		count++
		if count%k == 0 {
			//翻转链表[left,right]的部分变为[right,left]
			ReversePartOfList(left, right)
			//连接前面的部分
			if tail != nil {
				tail.Next = right
			}
			//新链表的头是第一次翻转后的right
			if count == k {
				head = right
			}
			//update
			tail = left
			right = left
			left = left.Next
		}
	}

	return head
}

//ReversePartOfList 翻转链表的[head...tail]部分。
func ReversePartOfList(head, tail *utility.ListNode) {
	var prev *utility.ListNode
	originHead := head
	guard := tail.Next
	for head != guard {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	originHead.Next = guard
}

//ReverseKGroupAdvanced 将链表的节点每k个一组进行翻转
func ReverseKGroupAdvanced(head *utility.ListNode, k int) *utility.ListNode {
	if k == 1 {
		return head
	}

	//哑巴节点
	dummy := &utility.ListNode{
		Val:  -1,
		Next: head,
	}
	prev, nHead := dummy, dummy.Next //要翻转的子区间的前驱节点和子区间的头节点
	count := 1                       //计数
	for head != nil {
		if count%k == 0 {
			//翻转[nHead,head]子区间。即从nHead.Next开始，每遍历一个节点，就将其交换到prev.Next
			for i := 1; i < k; i++ {
				//区间原地翻转，即将nHead.Next挂载到prev后面。其实就是依次把nHead后面的节点
				//挂载到prev后面。以0 1 2 3为例，nHead指向1
				//0 1 2 3 --> 0 2 1 3 --> 0 3 2 1。最后nHead就是区间的最后一个节点
				next := nHead.Next     //保存nHead.Next
				nHead.Next = next.Next //让nHead连接nHead.Next.Next
				next.Next = prev.Next  //原nHead.Next连接至prev后面
				prev.Next = next
			}
			//nHead此时指向当前区间的未元素，需要更新以指向下一个子区间的首元素
			prev = nHead
			nHead = nHead.Next
			head = nHead
		} else {
			head = head.Next
		}
		count++
	}

	return dummy.Next
}
