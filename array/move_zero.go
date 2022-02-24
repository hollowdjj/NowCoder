package array

/*
移动0
给定一个数组，请你实现将所有 0 移动到数组末尾并且不改变其他数字的相对顺序。
*/

func MoveZeros(nums []int) []int {
	left, right := 0, 0
	n := len(nums)
	for right < n {
		if nums[left] != 0 {
			left++
			right++
		} else {
			//当nums[left]为0时，那么right一直往后走找到left后面第一个不为0的数
			//然后交换nums[left]和nums[right]
			right = left + 1
			for right < n {
				if nums[right] != 0 {
					nums[left], nums[right] = nums[right], nums[left]
					break
				}
				right++
			}
			left++
		}
	}
	return nums
}
