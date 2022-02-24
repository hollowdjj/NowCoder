package array

/*
二分查找
请实现无重复数字的升序数组的二分查找

给定一个元素升序的、无重复数字的整型数组 nums 和一个目标值 target ，写一个
函数搜索 nums 中的 target，如果目标值存在返回下标（下标从 0 开始），否则返回 -1
*/

func BinarySearch(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			//说明target不可能在左数组，故往右数组找
			left = mid + 1
		} else if nums[mid] > target {
			//说明target不可能在右数组，故往左数组找
			right = mid - 1
		} else {
			return mid
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}
