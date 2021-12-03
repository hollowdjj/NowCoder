package array

/*
给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数
要求：空间复杂度 O(1)，时间复杂度O(logn)

二分查找复杂度分析：
假设有n个元素，第一次二分后，需要在n/2个元素中继续查找；第二次二分后，需要在n/2^2个元素
中进行查找；第t次二分后，需要在n/2^t个元素中进行查找。最坏的情况下，只剩下一个元素，也即
n/2^t = 1，那么有t = log(n)
*/

//GetNumberOfK 二分查找
func GetNumberOfK(data []int, k int) int {
	leftBound := binarySearch(data, k)
	rightBound := binarySearch(data, k+1)
	if leftBound == len(data) || data[leftBound] != k {
		return 0
	}

	return rightBound - leftBound
}

//binarySearch 二分法找有序数组中元素k第一次出现的位置。存在三种情况：
//1. 当k大于数组的最大元素时，返回值为len(array)
//2. 当k小于数组中的最小元素时，返回值为0;
//3. 当k大于数组中的最小元素，小于最大元素，但不存在于数组中时，返回的是第一个大于k的元素的位置
func binarySearch(array []int, k int) int {
	l, r := 0, len(array)
	for l < r {
		mid := l + (r-l)/2
		//当array[mid]大于等于k时，说明元素k首次出现的位置一定在[l,mid]之间
		//当array[mid]小于k时，说明元素k首次出现的位置一定在[mid+1,r]之间
		if array[mid] >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
