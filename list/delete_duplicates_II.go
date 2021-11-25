package list

import "nowcoder/utility"

/*
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
例如：
给出的链表为1→2→3→3→4→4→5, 返回1→2→5.
给出的链表为1→1→1→2→3, 返回2→3.

进阶：空间复杂度O(1)，时间复杂度O(n)
*/

func DeleteDuplicatesII(head *utility.ListNode) *utility.ListNode {
	dummy := &utility.ListNode{
		Val:  -1,
		Next: head,
	}
	//双指针
	prev, curr := dummy, head
	for curr != nil && curr.Next != nil {
		//如果当前值和下一个值相等，那么while循环找到下一个不相等的点，并将其挂在prev后面
		//否则curr和prev均往后移动一格
		if curr.Val == curr.Next.Val {
			temp := curr.Next
			for temp != nil && temp.Val == curr.Val {
				temp = temp.Next
			}
			prev.Next = temp
			curr = temp
		} else {
			prev = prev.Next
			curr = curr.Next
		}
	}
	return dummy.Next
}
