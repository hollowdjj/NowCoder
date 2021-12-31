package array

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007

数据范围：对于50% 的数据, size≤10^4；对于100%的数据, size≤10^5
数组中所有数字的值满足 0≤val≤1000000

要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
输入描述：题目保证输入的数组中没有的相同的数字

输入：[1,2,3,4,5,6,7,0]
输出：7

输入：[1,4,2,3,6,5,7,0]
输出：7 + 0 + 1 + 0 + 1 + 1 + 0 + 0 = 10

此题应该利用归并排序求解，即在归并两有序数组的适合统计逆序对的个数。因为，对两有序左右数组，例如[4 5] [1 2]
那么逆序对的个数显然为4个，即[4 1] [4 2] [5 1] [5 2]。也就是说，当左右数组都是有序数组且左数组中的第i个个数
大于右数组中的某个数时，left[i:]中的所有数对于右数组中的当前数而言都构成逆序对。
*/
var count = 0

func InversePairs(data []int) int {
	mergeSort(data, 0, len(data)-1)
	return count % 1000000007
}

//mergeSort归并排序
func mergeSort(array []int, left, right int) {
	mid := left + (right-left)/2
	if left < right {
		//递归合并左子数组
		mergeSort(array, left, mid)
		//递归合并右子数组
		mergeSort(array, mid+1, right)
		//此时，左右两子数组已经是有序的了，合并即可。
		mergeAndCount(array, left, mid, right)
	}
}

//mergeAndCount 合并两个有序数组并对逆序对计数
//@param left  左数组起始位置
//@param mid   左数组终止位置
//@param right 右数组终止位置
func mergeAndCount(array []int, left, mid, right int) {
	//生成一个临时数组，保存左右数组合并后的新数组。由于左右数组一定是紧挨着的
	//故新数组的元素个数为right-left+1
	temp := make([]int, right-left+1)

	//合并左右两有序数组
	index := 0
	l, r := left, mid+1
	for l <= mid && r <= right {
		if array[l] <= array[r] {
			temp[index] = array[l]
			l++
		} else {
			count += mid - l + 1
			temp[index] = array[r]
			r++
		}
		index++
	}

	for l <= mid {
		temp[index] = array[l]
		index++
		l++
	}
	for r <= right {
		temp[index] = array[r]
		index++
		r++
	}

	//将合并后的数组赋值给[left,right]
	for i := 0; i < len(temp); i++ {
		array[left] = temp[i]
		left++
	}
}
