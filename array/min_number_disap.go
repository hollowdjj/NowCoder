package array

/*
缺失的第一个正整数
给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数
例如：[1,0,2] 输出 3; [-2,3,4,1,5] 输出2
要求：空间复杂度O(1)，时间复杂度O(n)

*/

func MinNumberDisappeared(nums []int) int {
	//nums的长度为n，那么缺失的第一个正整数要么在[1,n]中，要么就是n+1(即nums为1,2,3,4,5...n)
	//因此，我们首先遍历一遍数组，将所有的负数更改为n+1。然后再遍历一下数组，将数组中所有大于等
	//于1且小于等于n的数修改为负数，也就是将那些在[1,n]中出现过的数做一个标记。最后，再遍历一次
	//数组，第一个非负数的下标加1即是结果。这是一个很重要的自身hash方法。
	n := len(nums)
	//遍历数组，将所有负数都修改为n+1。这一步的目的主要是剔除负数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	//再次遍历数组。若1<=num[i]<=n，那么就将nums[nums[i]]变为负数。相当于标记[1,n]中的
	//数x=num[i]已存在
	for i := 0; i < n; i++ {
		if num := nums[i]; AbsInt(num) <= n {
			nums[AbsInt(num)-1] *= -1
		}
	}
	//再次遍历数组，遇到的第一个非负数数的下标加1即为答案
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
