package jz

/*
给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数

输入：[1,2,3,3,3,3,4,5],3
输出：4
*/

func GetNumberOfK(data []int, k int) int {
	//二分法找第一个小于k的元素，和第一个大于k的元素
	i, j := firstSmaller(data, k), firstLarger(data, k)
	if i == j {
		return 0
	}
	return j - i - 1
}

//二分找到第一个小于target的元素。
//当所有元素都大于target或者第一个元素为target时，返回-1；
//当所有元素都小于target时，返回n
func firstSmaller(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

//二分法找第一个大于target的元素。
//当所有元素都大于target时，返回-1；
//当所有元素都小于target时，返回n
func firstLarger(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
