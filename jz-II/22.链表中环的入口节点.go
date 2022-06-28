package jz_II

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	//先判断有无环
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	//将slow移动到开头，快慢指针每次都移动一步，相遇处为入口节点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
