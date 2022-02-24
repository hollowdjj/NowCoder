package array

/*
最小的K个数
给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组
元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。
要求：空间复杂度 O(n)O(n) ，时间复杂度 O(nlogn)O(nlogn)
*/

func GetLeastNumbers_Solution(input []int, k int) []int {
	//归并排序数组
	n := len(input)
	if n < k {
		return nil
	}
	mergeSort1(input, 0, len(input)-1)
	return input[0:k]
}

func mergeSort1(arr []int, left, right int) {
	if left < right {
		mid := (right + left) / 2
		mergeSort1(arr, left, mid)
		mergeSort1(arr, mid+1, right)
		merge1(arr, left, mid, right)
	}
}

//合并两个已排序好的数组
func merge1(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j := left, mid+1
	k := 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	//赋值给原数组
	k--
	for ; k >= 0; k-- {
		arr[right] = temp[k]
		right--
	}
}
