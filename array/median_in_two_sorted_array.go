package array

/*
在两个长度相等的升序数组中找到上中位数
给定两个递增数组arr1和arr2，已知两个数组的长度都为N，求两个数组中所有数的上中位数。
上中位数：假设递增序列长度为n，为第n/2个数
例如，输入：[1,4,6,7,9] [2,3,4,5,8], 输出：

要求：时间复杂度 O(n)，空间复杂度 O(1)
进阶：时间复杂度为O(logN)，空间复杂度为O(1)
*/

//FindMedianInTwoSortedArray 双指针法，时间复杂度为O(n)
func FindMedianInTwoSortedArray(arr1, arr2 []int) int {
	/*
		基本思路为：首先计算出当两个数组合并成一个数组时，上中位数所在的位置。然后，使用
		两个指针分别指向两个数组的头部，开始遍历。
	*/
	n1, n2 := len(arr1), len(arr2)
	mid := (n1 + n2) / 2 //上中位数的位置，即第mid个数
	p1, p2 := 0, 0
	temp := 0
	for p1 < n1 && p2 < n2 {
		if arr1[p1] < arr2[p2] {
			temp = arr1[p1]
			p1++
		} else {
			temp = arr2[p2]
			p2++
		}
		if p1+p2 == mid {
			break
		}
	}

	return temp
}

//FindMedianInTwoSortedArrayAdvanced 二分法，时间复杂度为O(logN)
func FindMedianInTwoSortedArrayAdvanced(arr1, arr2 []int) int {
	/*
		当要求复杂度为O(logN)时，应该相当是使用二分法。此时，需要注意区分长度
		为奇数还是偶数两种情况。(注意，在牛客网中，说第n个数的时候，一般是以1开头的)。
		此时，若n为偶数，例如：[1,2,3,4][4,5,6,7]：
		1. arr1[mid1] < arr2[mid2]，说明arr1[0:mid1]中的所有数是小于arr2[mid2:right]的。
		因此，当合并成一个大数组时，这两部分应分别位于数组的最左边和最右边，上中位数一定在arr1[mid1+1,right]
		和arr2[0:mid2]中（因为，此时最左和最右两部分长度是相等的）
		2. arr1[mid1] > arr2[mid2]。同理，上中位数一定在arr1[0:mid1]和arr2[mid2+1:right]
		3. arr1[mid1] == arr2[mid2] 则上中位数就是两数组的上中位数
		此时，若n为奇数：
		1. arr1[mid1] < arr2[mid2]，上中位数在arr1[mid1+1:right]和arr2[0:mid2+1]
		2. arr1[mid1] > arr2[mid2]，上中位数在arr2[mid2+1:right]和arr1[o:mid1+1]
	*/
	return 1
}
