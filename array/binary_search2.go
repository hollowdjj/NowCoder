package array

/*
二分查找(二)

请实现有重复数字的升序数组的二分查找
给定一个元素有序的（升序）长度为n的整型数组 nums 和一个目标值 target  ，写一个函数
搜索 nums 中的第一个出现的target，如果目标值存在返回下标，否则返回 -1
*/

func BinarySearch2(nums []int, target int) int {
	//思路就是找到第一个小于target的元素，他的下一个元素就是target第一次出现的位置
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if nums[right+1] == target {
		return right + 1
	}
	return -1
}
