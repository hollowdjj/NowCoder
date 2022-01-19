package others

import "math"

/*
给定一个整数数组，你需要找出一个连续子数组，将这个子数组升序排列后整个数组都将是升序数组。
请你找出满足题设的最短的子数组。

输入:[2,6,4,8,10,9,15]
输出:5 (只需对6,4,8,10,9进行重新排序即可)
*/

func FindUnsortedSubArray(nums []int) int {
	//在一个有序数组nums中，nums[i]一定是nums[i:]中的最小值，也一定是nums[:i]中的
	//最大值。因此，我们可以利用这个性质进行求解。首先，从左往右遍历，找到最后一个
	//不是nums[:i]中最大值的i，作为右边界。然后，从右往左遍历，找到最后一个不是nums[i:]
	//中的最小值的j，作为左边界。
	n := len(nums)
	if n == 1 {
		return 0
	}
	lb, rb := -1, -1
	min, max := math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		//从左往右遍历，确定右边界
		if nums[i] >= min {
			min = nums[i]
		} else {
			rb = i
		}
		//从右往左遍历，确定左边界
		if nums[n-1-i] <= max {
			max = nums[n-1-i]
		} else {
			lb = n - 1 - i
		}
	}
	if lb == -1 {
		return 0
	}
	return rb - lb + 1
}
