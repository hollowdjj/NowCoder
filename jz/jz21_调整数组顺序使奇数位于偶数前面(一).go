package jz

/*
输入一个长度为 n 整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前面部分，所有的偶数位于数组的后面
部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

输入：[2,4,6,5,7]
输出：[5,7,2,4,6]
*/

func reOrderArray(array []int) []int {
	n := len(array)
	i := 0 //奇数序列最后一个数的下一个位置
	for j := 0; j < n; j++ {
		if array[j]%2 == 1 {
			temp := array[j]
			//将[i,j-1]的数字往后移动一格
			for k := j - 1; k >= i; k-- {
				array[k+1] = array[k]
			}
			array[i] = temp
			i++
		}
	}
	return array
}
