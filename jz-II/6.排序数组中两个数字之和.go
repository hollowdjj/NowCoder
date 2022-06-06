package jz_II

func TwoSum(numbers []int, target int) []int {
	n := len(numbers)
	res := make([]int, 2)
	for i := 0; i < n; i++ {
		if numbers[i] >= target && target > 0 {
			break
		}
		t := find(numbers, target-numbers[i])
		if t != -1 && i != t {
			res[0], res[1] = i, t
			break
		}
		//形如[1,3,4,4,5,6],8
		if t == i && i < n-1 && numbers[i+1] == numbers[i] {
			res[0], res[1] = i, i+1
			break
		}
	}
	return res
}

func find(nums []int, target int) int {
	//二分法找nums中target的下标，若不存在，返回-1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
