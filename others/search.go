package others

/*
在旋转过的有序数组中寻找目标值
有一个长度为 n 的按严格升序排列的整数数组 nums ，在实行 search 函数之前，在某个下标 k 上进行旋转，使数组
变为[nums[k],nums[k+1],.....,nums[nums.length-1],nums[0],nums[1],.......,nums[k-1]]。
给定旋转后的数组 nums 和一个整型 target ，请你查找 target 是否存在于 nums 数组中并返回其下标（从0开始计数），如果不存在请返回-1。


比如，数组[0,2,4,6,8,10]在下标3处旋转之后变为[6,8,10,0,2,4], 当给定target为10时，10的下标是2，target为3时，nums数组中不存在3，所以返回-1
*/

func Search(nums []int, target int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		if target == nums[i-1] {
			return i - 1
		}
		if nums[i] < nums[i-1] {
			if target > nums[i-1] || target > nums[n-1] || target < nums[i] {
				//此时nums[i]为数组的最小值，nums[i-1]为数组的最大值
				return -1
			}
		}
	}
	if nums[n-1] == target {
		return n - 1
	}
	return -1
}
