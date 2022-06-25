package jz_II

import "sort"

/*
给定一个包含n个整数的数组nums，判断nums中是否存在三个元素a ，b ，c ，使得a + b + c = 0 ？
请找出所有和为0且不重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if nums[i]+sum > 0 {
				right--
			} else if nums[i]+sum < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//去重
				b, c := nums[left], nums[right]
				for left < right && nums[left] == b {
					left++
				}
				for left < right && nums[right] == c {
					right--
				}
			}
		}
	}
	return res
}
