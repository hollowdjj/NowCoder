package list

import (
	"nowcoder/utility"
)

/*
将一个节点数为size链表m位置到n位置之间的区间反转。例如：
给出的链表为1→2→3→4→5→NULL, m=2,n=4m=2,n=4,返回1→4→3→2→5→NULL。其中0<m≤n≤size

此题是关于局部链表翻转的，之前遇到的时候都是采用翻转再链接的方式，不太优雅，这里的解法
需要牢牢记住。
*/

func ReverseBetween(head *utility.ListNode, m int, n int) *utility.ListNode {
	//哑巴节点
	dummy := &utility.ListNode{
		Val:  -1,
		Next: head,
	}
	//找到第m-1个节点
	guard := dummy
	for i := 1; i < m; i++ {
		guard = guard.Next
	}
	//翻转第m个节点到第n个节点间的所有节点(从第m+1个节点开始，每遍历到一个节点就将他换到子区间的头部)
	head = guard.Next        //第m个节点
	for i := m; i < n; i++ { //翻转
		//将head.Next交换到guard.Next
		next := head.Next      //head.Next
		head.Next = next.Next  //断开head.Next，链接head和head.Next.Next
		next.Next = guard.Next //将head.Next链接至guard.Next
		guard.Next = next
	}

	return dummy.Next
}
