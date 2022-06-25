package jz_II

/*
给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	n := len(nums)
	res := 0
	mul := 1
	for right < n {
		//滑动窗口，当子数组ABCD的乘积小与k时，我们取D,CD,BCD,ABCD都是
		//满足题意的。故res += right - left + 1。这里必须取以D，也就是
		//最后一个元素结尾的连续子数组，否则会出现重复。例如取BC，会和只有
		//ABC的时候取BC重复。
		mul *= nums[right]
		for left <= right && mul >= k {
			mul /= nums[left]
			left++
		}
		res += right - left + 1
		right++
	}
	return res
}
