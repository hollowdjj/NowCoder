package array

//标准二分。
func BinarySearch(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1
	//这里是left<=right。因为使用的是闭区间[left,right]，left == right是有意义的
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] < target {
			//说明target不可能在左数组，故往右数组找
			left = mid + 1
		} else if nums[mid] > target {
			//说明target不可能在右数组，故往左数组找
			right = mid - 1
		} else {
			return mid
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

//二分查找在nums中找到第一个小于target的数
func BinarySearchFirstSmaller(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		//因为是要找到第一个比target小的数。因此，当nums[mid] == target时，我们需要缩小右边界
		//即让right = mid - 1。所以这里合并成了nums[mid] >= target
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}

//二分法在nums找到第一个大于target的数
func BinarySearchFirstBigger(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		//因为是要找到第一个大于target的数。因此，当nums[mid] == target时，需要增大左边界。
		//即left = mid + 1。所以这里合并成了nums[mid]<=target
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[left]
}
