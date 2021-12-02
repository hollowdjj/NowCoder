package list

import "nowcoder/utility"

//判断链表是否有环，要求空间复杂度O(1)时间复杂度O(n)

//HasCycleFast 速度较快的解法，但缺点在于会修改节点中的值，循环次数=2*环长 + 链表非环部分的长度
func HasCycleFast(head *utility.ListNode) bool {
	for {
		if head == nil {
			return false
		}
		if head.Val > 100000 {
			//出现大于100000的值，说明存在环，此时还原节点值并返回true
			temp := head
			for {
				head.Val -= 200001
				head = head.Next
				if head == temp {
				}
				return true
			}
		} else {
			head.Val += 200001
		}
		head = head.Next
	}
}

/*
HasCycleClassic 经典快慢指针解法。快慢指针会在环中的某一点相遇。因为一旦有环，就可以视为
快指针在追赶慢指针，而快指针每次多走一步，因此一定能追上。
*/
func HasCycleClassic(head *utility.ListNode) bool {
	slow, fast := head, head

	for {
		if fast == nil {
			return false
		}

		//快指针移动两步
		if temp := fast.Next; temp != nil {
			fast = temp.Next
		} else {
			return false
		}
		//慢指针移动一步
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
}
