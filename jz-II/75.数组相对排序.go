package jz_II

/*
给定两个数组，arr1和arr2，

arr2中的元素各不相同, arr2中的每个元素都出现在arr1中
对arr1中的元素进行排序，使arr1中项的相对顺序和arr2中的相对顺序相同。
未在arr2中出现过的元素需要按照升序放在arr1的末尾。

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]
*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	//采用计数排序，数组arr[i]表示数字i出现的次数。
	//首先遍历arr1，得到最大值max，并初始化arr，使len(arr)=max
	max := -1
	for _, v := range arr1 {
		if v > max {
			max = v
		}
	}
	//再次遍历arr1，构建arr
	arr := make([]int, max+1)
	for _, v := range arr1 {
		arr[v]++
	}
	//遍历arr2，填充结果数组中的前半部分
	k := 0
	res := make([]int, len(arr1))
	for _, v := range arr2 {
		for i := 0; i < arr[v]; i++ {
			res[k] = v
			k++
		}
		arr[v] = 0
	}
	//遍历arr。此时若arr[i] != 0，表示i没在arr2中出现过
	for i, v := range arr {
		for j := 0; j < v; j++ {
			res[k] = i
			k++
		}
	}
	return res
}
