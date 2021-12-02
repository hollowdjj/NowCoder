package list

import "nowcoder/utility"

/*
输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。
要求：空间复杂度 O(n)，时间复杂度 O(n)
进阶：空间复杂度 O(1)，时间复杂度 O(n)
*/

//FindKthToTail 使用哈希表辅助找到链表倒数第k个节点。若k大于链表元素个数，返回nil
func FindKthToTail(pHead *utility.ListNode, k int) *utility.ListNode {
	m := make(map[int]*utility.ListNode)
	count, i := 0, 0
	for pHead != nil {
		i++
		m[i] = pHead
		count++
		pHead = pHead.Next
	}

	return m[count-k+1] //一共有n个数，倒数第k个数就是正数第n-k+1个数
}

//FindKthToTailAdvanced 快慢指针找到链表倒数第K个节点。若k大于链表元素个数，返回nil
func FindKthToTailAdvanced(pHead *utility.ListNode, k int) *utility.ListNode {
	//快慢指针。快指针先向前走k步，然后两指针在一起每次走一步
	slow, fast := pHead, pHead
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
