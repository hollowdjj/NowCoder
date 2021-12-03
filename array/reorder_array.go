package array

/*
输入一个长度为 n 整数数组，数组里面不含有相同的元素，实现一个函数来调整该数组中数字的顺序，
使得所有的奇数位于数组的前面部分，所有的偶数位于数组的后面部分，并保证奇数和奇数，偶数和偶
数之间的相对位置不变。
要求：时间复杂度 O(n)，空间复杂度O(n)
进阶：时间复杂度 O(n^2)，空间复杂度O(1)
*/

//ReorderArray 使用额外数组辅助进行奇偶重排
func ReorderArray(array []int) []int {
	var odd, even []int
	for _, v := range array {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	odd = append(odd, even...)
	return odd
}

//ReorderArrayAdvanced 原地进行奇偶重排
func ReorderArrayAdvanced(array []int) []int {
	i := 0
	for j := 0; j < len(array); j++ {
		//如果array[j]为奇数，那么将[i,j-1]元素向后移动，并将j元素赋值给i
		if array[j]%2 == 1 {
			temp := array[j]
			//
			for k := j - 1; k >= i; k-- {
				array[k+1] = array[k]
			}
			array[i] = temp
			i++
		}
	}
	return array
}
