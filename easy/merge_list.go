package easy

import "fmt"

/*
输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
要求：空间复杂度 O(1)，时间复杂度 O(n)
如输入{1,5,9},{2,3,4,7}时，合并后的链表为{1,2,3,4,5,7,9}
*/

func MergeList(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	start := &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := start //哨兵节点
	for pHead1 == nil && pHead2 == nil {
		if pHead1 != nil && (pHead2 == nil || pHead1.Val <= pHead2.Val) {
			temp.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			temp.Next = pHead2
			pHead2 = pHead2.Next
		}
		temp = temp.Next
	}

	return start.Next
}

func TestMergeList() {
	PrintList(MergeList(GenList(nil), GenList([]int{2, 3, 4, 7})))
}

func GenList(slice []int) *ListNode {
	n := len(slice)
	if n < 1 {
		return nil
	}
	pHead := &ListNode{
		Val:  slice[0],
		Next: nil,
	}
	temp := pHead
	for i := 1; i < n; i++ {
		temp.Next = &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		temp = temp.Next
	}
	return pHead
}

func PrintList(pHead *ListNode) {
	fmt.Print("[")
	for pHead != nil {
		fmt.Printf(" %d ", pHead.Val)
		pHead = pHead.Next
	}
	fmt.Println("]")
}
