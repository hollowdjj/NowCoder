package easy

/*
删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
例如：
给出的链表为1->1->2 返回1->2.
给出的链表为1→1→2→3→3,返回1→2→3.

数据范围：链表长度满足 0≤n≤100，链表中任意节点的值满足∣val∣≤100
进阶：空间复杂度 O(1)，时间复杂度 O(n)
*/

func deleteDuplicates(head *ListNode) *ListNode {
	//双指针法，left和right分别表示重复序列的第一个元素和最后一个元素的后一个元素
	if head == nil {
		return nil
	}

	left, right := head, head.Next
	for right != nil {
		if left.Val != right.Val {
			left = left.Next
		}
		right = right.Next
		left.Next = right
	}

	return head
}

func TestDeleteDuplicates() {
	head := GenList([]int{1, 1, 2})
	PrintList(deleteDuplicates(head))
}
