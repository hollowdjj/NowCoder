package array

/*
二分查找(二)

请实现有重复数字的升序数组的二分查找
给定一个元素有序的（升序）长度为n的整型数组 nums 和一个目标值 target  ，写一个函数
搜索 nums 中的第一个出现的target，如果目标值存在返回下标，否则返回 -1
*/

func BinarySearch2(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left < right {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			//往mid的左边进行搜索，找到第一个目标值出现的位置
			for i := mid; i >= 0; i-- {
				if nums[i] != target {
					return i + 1
				}
			}
			return 0
		}
	}

	if nums[left] != target {
		//没找到返回-1
		return -1
	}
	//对于[1,2,3]找3的情况
	return left
}
