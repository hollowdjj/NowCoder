package list

import (
	"fmt"
	"nowcoder/utility"
)

/*
编号为 1 到 n 的 n 个人围成一圈。从编号为 1 的人开始报数，报到 m 的人离开。下一个人继续从 1 开始报数。
n-1 轮结束以后，只剩下一个人，问最后留下的这个人编号是多少？
*/

//Ysf 动态规划解约瑟夫环问题
func Ysf(m, n int) int {
	/*
		第一个出列的人的序号为k-1 = (m-1)%n，下一个开始报数的人的序号为k = m%n。随后，剩下的n-1个人组成了
		一个新的约瑟夫环k k+1 k+2...n-2 n-1 0 1 2...k-3 k-2(k-1出列了故最多只能到k-2)。然后，对这个约瑟
		夫环的编号做一下转换：
						x':	k k+1 k+2...n-2 n-1 0 1 2...k-3 k-2
		                x : 0  1   2      .....         n-3 n-2
		对应关系为x' = (x+k) % n。也就是说，假设x为n-1个人的约瑟夫环问题的解，那么n个人的解为(x+k) % n。
		设dp[i]为i个人的解，那么dp[i] = (dp[i-1]+k)%i = (dp[i-1]+m%i)%i = (dp[i-1]+m)%i
	*/
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + m) % i
	}
	return dp[n] + 1
}

//YsfByList 环形链表的约瑟夫环问题
func YsfByList(head *utility.ListNode, m int) *utility.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 1
	prev := head
	temp := head
	for {
		if count%m == 0 {
			//删除当前链表节点
			next := head.Next
			if next == head {
				break
			}
			prev.Next = next
			head = next
			prev = head
			count++
		}
		count++
		prev = head
		head = head.Next
	}
	fmt.Println(temp.Val)
	return head
}
