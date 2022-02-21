package array

/*
数组排序
*/

//冒泡排序。最优O(n)，最坏O(n^2)
func BubbleSort(arr []int) []int {
	//从左到右依次比较arr[j]和arr[j-1]的大小，并将大的那个交换到右边。一趟交换完成后，数组中的最后一个
	//元素即为数组的最大值。第二趟交换完成后，倒数第二个元素为数组倒数第二大的值。因此，n个元素需要进行
	//n-1趟交换。而第i趟交换又要进行n-i次左右交换。因此，冒泡排序的最优时间复杂度为O(n)，即只需要一趟
	//交换即可。最坏情况下的时间复杂度为O(n^2)，即数组本身是逆序的(4,3,2,1,0)。
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

//快排，填坑式做法。最优O(nlogn)，最坏O(n^2)
func QuickSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	//需要注意的是，我们必须选择数组的第一个数作为pivot，而不能选择其他的数。一些改进方法，例如随机选取
	//数组中的一个数做为pivot时，也只是将那个随机选取的数和第一个数交换，pivot始终是指数组的头元素。
	pivot := arr[0]
	left, right := 0, len(arr)-1
	for left < right {
		//首先从右往左找到第一个小于pivot的数。然后将其赋值给arr[left](之前挖好的坑)
		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			left++
		}
		//然后从左往右找到第一个大于pivot的数。然后将其赋值给arr[right](之前挖好的坑)
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right--
		}
	}
	arr[left] = pivot
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}
