package array

/*
旋转数组的最小数字

有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，即把一个数组最开始的若干个元素搬到
数组的末尾，变成一个旋转数组，比如变成了[3,4,5,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个
旋转数组，求数组中的最小值。

数据范围：1≤n≤10000，数组中任意元素的值: 0≤val≤10000
要求：空间复杂度：O(1)，时间复杂度：O(logn)
*/

func MinNumberInRotateArray(rotateArray []int) int {
	//最基础的方法就是遍历数组，找到交换的地方。最坏的情况下时间复杂度为O(n)
	n := len(rotateArray)
	for i := 0; i < n-1; i++ {
		if rotateArray[i] < rotateArray[i-1] {
			return rotateArray[i]
		}
	}
	return rotateArray[0]
}

func MinNumberInRotateArrayAdvanced(rotateArray []int) int {
	left, right := 0, len(rotateArray)-1
	for left < right {
		mid := left + (right-left)/2
		if rotateArray[mid] < rotateArray[right] {
			//说明右数组时有序的，此时往左数组找
			right = mid
		} else if rotateArray[mid] > rotateArray[right] {
			//说明右数组无序，往右数组找
			left = mid + 1
		} else {
			//rotateArray[mid] == rotateArray[right] [3,4,2,2,2]
			right--
		}
	}
	return rotateArray[left]
}
