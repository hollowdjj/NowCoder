package jz

/*
输入一个升序数组 array 和一个数字S，在数组中查找两个数，使得他们的和正好是S，如果有多对数字的和等于S，返回任意一组即可，如果
无法找出这样的数字，返回一个空数组即可。

输入：[1,2,4,7,11,15],15
输出：[4,11] or [11,4]
*/

func FindNumbersWithSum(array []int, sum int) []int {
	//双指针
	n := len(array)
	left, right := 0, n-1
	for left < right {
		temp := array[left] + array[right]
		if temp == sum {
			return []int{array[left], array[right]}
		} else if temp > sum {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
