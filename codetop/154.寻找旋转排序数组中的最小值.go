package codetop

/*
已知一个长度为n的数组，预先按照升序排列，经由1到n次旋转后，得到输入数组。例如，原数组
nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组[a[0], a[1], a[2], ..., a[n-1]]旋转一次 的结果为数组[a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返
回数组中的 最小元素。你必须尽可能减少整个过程的操作步骤。

输入：nums = [2,2,2,0,1]
输出：0
*/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] < nums[right] {
			//nums[mid...right]有序，需要进入无序的部分查找，故right=mid
			right = mid
		} else if nums[mid] > nums[right] {
			//nums[mid...right]本身就是无序的，因此缩小范围使left=mid+1
			left = mid + 1
		} else {
			/*
				若nums[right]是唯一最小值，那就不可能满足判断条件nums[mid] == nums[right]
				若nums[right]不是唯一最小值，由于mid<right，而nums[mid] == nums[right]
				即还有最小值存在于[left,right−1]区间，因此不会丢失最小值。
			*/
			right--
		}
	}
	return nums[left]
}
