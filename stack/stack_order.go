package stack

import "nowcoder/utility"

/*
给你一个 1 到 n 的排列和一个栈，并按照排列顺序入栈

你要在不打乱入栈顺序的情况下，仅利用入栈和出栈两种操作，输出字典序最大的出栈序列

排列：指 1 到 n 每个数字出现且仅出现一次

数据范围：  1≤n≤5×10^4，排列中的值都满足1≤val≤n

进阶：空间复杂度O(n)，时间复杂度O(n)
*/

func Order(a []int) []int {
	//要得到字典序最大的出栈序列，那么当前要出栈的元素a[i]一定要比之后将要进栈的元素a[i+1:]都大
	//因此，一个空间复杂度为O(n)，时间复杂度为O(n)的方法为：
	//首先，使用一个辅助数组dp。其中dp[i]表示子数组a[i+1:]中的最大值。
	//遍历数组a。若栈为空，那么a[i]入栈；若栈不为空，且a[i]大于等于dp[i]，那么a[i]不入栈，直接添加到返回数组中
	//数组a遍历完成后，若栈不为空，那么依次pop到返回数组中即可。

	//从后往前遍历数组a，得到辅助数组dp。其中状态转移方程为：dp[i] = Max{dp[i+1],a[i+1]}
	n := len(a)
	dp := make([]int, n)
	dp[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = max(a[i+1], dp[i+1])
	}

	var s utility.Stack
	res := make([]int, n)
	j := 0
	for i, _ := range a {
		if a[i] > dp[i] {
			res[j] = a[i]
			j++
			//这里需要特别注意的是，我们还需检查当前栈顶元素是否大于之后将要入栈的所有元素的最大值。
			//例如，[2,6,7,5,4,3]。当遍历到元素7时，7被加入res中。此时6大于7后面的所有数，6也需要
			//出栈。这里也解释了为什么dp[i]表示的是子数组a[i+1:]中的最大值。
			for !s.Empty() {
				top := (*s.Top()).(int)
				if top > dp[i] {
					res[j] = (*s.Pop()).(int)
					j++
				} else {
					break
				}
			}
		} else {
			s.Push(a[i])
		}
	}
	for !s.Empty() {
		res[j] = (*s.Pop()).(int)
		j++
	}
	return res
}
