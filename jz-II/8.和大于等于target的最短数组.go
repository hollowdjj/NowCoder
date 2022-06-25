package jz_II

import "math"

/*
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的连续子数组[numsl, numsl+1, ..., numsr-1, numsr]
并返回其长度。如果不存在符合条件的子数组，返回0。

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	n := len(nums)
	res := math.MaxInt32
	for right < n {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < res {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if res == math.MaxInt32 {
		res = 0
	}
	return res
}
