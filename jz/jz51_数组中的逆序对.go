package jz

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。
并将P对1000000007取模的结果输出。 即输出P mod 1000000007

输入：[1,2,3,4,5,6,7,0]
输出：7
*/

var res51 = 0

func InversePairs(nums []int) int {
	merge(nums, 0, len(nums)-1)
	return res51
}

func merge(nums []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		merge(nums, left, mid)
		merge(nums, mid+1, right)
		mergeTwo(nums, left, mid, right)
	}
}

func mergeTwo(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			//例如[3,4] [1,2] 由于两个都是升序数组，所以在左数组中，从i往后的所有数字都可以组成一个逆序对
			res51 += mid - i + 1
			res51 %= 1000000007
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}

	for i = 0; i < len(temp); i++ {
		nums[left] = temp[i]
		left++
	}
}
