package list

import "nowcoder/utility"

/*
题目：
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
例如，给出的链表为1→2→3→3→4→4→5, 返回1→2→5；给出的链表为1→1→1→2→3, 返回2→3.
要求空间复杂度为O(1)，时间复杂度为O(n)

思路：
这道题和DeleteDuplicates很像，不过后者只是删除重复元素，使得所有元素都只出现一次，而本题
则是要删除所有重复的元素，只保留链表中仅出现过一次的元素，思路很简单，就是使用三个指针。
*/

//DeleteDuplicatesIIV1 删除有序链表中的所有重复元素，循环查找的版本
func DeleteDuplicatesIIV1(head *utility.ListNode) *utility.ListNode {
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

//DeleteDuplicatesIIV2 删除有序链表中的所有重复元素，三指针版本
func DeleteDuplicatesIIV2(head *utility.ListNode) *utility.ListNode {
	if head == nil {
		return nil
	}
	dummy := &utility.ListNode{Val: -1, Next: head}
	//left和right双指针用于寻找相邻重复元素，prev为left的前一个元素
	prev, left, right := dummy, head, head.Next
	for right != nil {
		if left.Val != right.Val {
			/*
				当left和right的值不相等时，prev和left指针移动。这里需要注意的是，我们不能
				让prev = prev.Next，否则在消除了一个重复序列后，后续操作中prev和left将始终
				指向同一个元素。
			*/
			prev = left
			left = left.Next
		} else {
			prev.Next = right.Next
			left = prev
		}
		right = right.Next
	}

	return dummy.Next
}
