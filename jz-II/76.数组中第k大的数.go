package jz_II

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
*/

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		adjust76(nums, i, n)
	}

	i := n - 1
	for k > 0 {
		nums[0], nums[i] = nums[i], nums[0]
		adjust76(nums, 0, i)
		k--
		i--
	}
	return nums[i+1]
}

func adjust76(nums []int, parent, end int) {
	child := 2*parent + 1
	for child < end {
		if child+1 < end && nums[child+1] > nums[child] {
			child++
		}
		if nums[parent] >= nums[child] {
			break
		}
		nums[parent], nums[child] = nums[child], nums[parent]
		parent = child
		child = 2*parent + 1
	}
}
