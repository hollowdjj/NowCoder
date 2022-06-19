package jz_II

/*
求山峰数组的山顶元素，例如：
输入 arr = [1,3,5,4,2]
输出 2
*/

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] < arr[mid+1] {
			//山峰一定在右边
			left = mid + 1
		} else {
			//山峰一定在左边
			right = mid
		}
	}
	return left
}
