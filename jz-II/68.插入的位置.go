package jz_II

/*
给定一个排序的整数数组 nums和一个整数目标值target，请在数组中找到target，并返回其下标。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/

func searchInsert(nums []int, target int) int {
	//找到有序数组nums中第一个大于target的数的索引
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	//需要找到最后一个小于target的数的位置
	if left > 0 && nums[left-1] == target {
		left, right = 0, left-2
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			}
		}
	}
	return left
}
