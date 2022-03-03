package list

/*
对链表执行插入排序

对链表进行插入排序，
从链表第一个元素开始可以视为部分已排序，每次操作从链表中移除一个元素，然后原地将移除的元素插入到已排好序的部分。
数据范围：链表长度满足1≤n≤1000，链表中每个元素满足∣val∣≤10000
2---------->4---------->1
			|
1---------->2---------->4
*/

func InsertSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: -10001} //已排序好序列的队列头
	tail := dummy                   //已排序好序列的队列尾
	curr := head
	for curr != nil {
		if curr.Val >= tail.Val {
			//直接插入已排序好序列的尾部
			tail.Next = curr
			tail = curr
			curr = curr.Next
			//与原链表断开
			tail.Next = nil
		} else {
			next := curr.Next
			//从已排序好序列的队列头开始，找到第一个节点值大于curr节点值的节点，然后将curr插入到他前面
			prev := dummy
			start := dummy.Next
			for start.Val < curr.Val {
				prev = start
				start = start.Next
			}
			prev.Next = curr
			curr.Next = start
			curr = next
		}
	}
	return dummy.Next
}
