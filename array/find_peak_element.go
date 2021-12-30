package array

/*
给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。
1.峰值元素是指其值严格大于左右相邻值的元素。严格大于即不能有等于
2.假设 nums[-1] = nums[n] = −∞
3.对于所有有效的 i 都有 nums[i] != nums[i + 1]
4.你可以使用O(logN)的时间复杂度实现此问题吗？
*/

func FindPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	return 0
}

func FindPeakElementAdvanced(nums []int) int {
	//核心思想在于：下坡不一定能找到解，但是上坡一定能找到解(最坏就是边界处)
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			//说明右边的是下坡，不一定有坡峰
			right = mid
		} else {
			//说明左边是上坡
			left = mid + 1
		}
	}

	return right
}
