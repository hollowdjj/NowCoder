package array

/*
旋转数组
一个数组A中存有n个整数，在不允许使用额外数组的情况下，将每个整数循环向右移动M个位置。
数据范围：0<n≤10, 0≤m≤1000
要求：空间复杂度O(1)，时间复杂度O(n)
*/

//RotateArray 翻转数组的方法完全旋转数组
func RotateArray(a []int, m, n int) []int {
	if n == 0 {
		return a
	}
	//相当于将倒数第index个元素及其之后的元素都移动到数组的头部
	index := m % n
	//首先翻转整个数组
	ReverseArray(a)
	//然后翻转[0,index-1]
	ReverseArray(a[0:index])
	//最后再翻转[index:]
	ReverseArray(a[index:])

	return a
}

//ReverseArray 双指针原地翻转数组
func ReverseArray(array []int) {
	n := len(array)
	i, j := 0, n-1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
}
