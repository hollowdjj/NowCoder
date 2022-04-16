package jz

/*'
有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，即把一个数组最开始的若干个元素搬到数组的末尾，变成一个
旋转数组，比如变成了[3,4,5,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。
*/

func minNumberInRotateArray(rotateArray []int) int {
	n := len(rotateArray)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if rotateArray[mid] > rotateArray[right] {
			//说明旋转点一定在[mid+1,right]中
			left = mid + 1
		} else if rotateArray[mid] < rotateArray[right] {
			//说明旋转点一定在[left,mid]中
			right = mid
		} else {
			//无法判断旋转点在那里，于是right--。例如[2,2,2,2,1,2]
			right--
		}
	}
	return rotateArray[left]
}
