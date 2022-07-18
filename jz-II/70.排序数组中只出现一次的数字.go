package jz_II

/*
给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
*/

func singleNonDuplicate(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
