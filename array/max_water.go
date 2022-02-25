package array

/*
接雨水问题
给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个柱
子高度图，计算按此排列的柱子，下雨之后能接多少雨水。(数组以外的区域高度视为0)
*/

func MaxWater(arr []int) int64 {
	n := len(arr)
	if n < 3 {
		return 0
	}

	//首先遍历一遍数组找到最高的柱子
	top := 0
	for i := 1; i < n; i++ {
		if arr[i] > arr[top] {
			top = i
		}
	}

	water := 0
	//在最高柱子的左边，我们使用双指针从左往右遍历。当arr[left] > arr[right]时，才
	//可盛水。否则让left = right
	left, right := 0, 0
	for right < top {
		if arr[left] > arr[right] {
			water += arr[left] - arr[right]
		} else {
			left = right
		}
		right++
	}

	//在最高柱子的右边，我们使用双指针从右往左遍历。当arr[left] < arr[right]时，才
	//可以盛水。否则让right = left
	left, right = n-2, n-1
	for left > top {
		if arr[left] < arr[right] {
			water += arr[right] - arr[left]
		} else {
			right = left
		}
		left--
	}

	return int64(water)
}
