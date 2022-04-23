package jz

/*
输入一个长度为 n 整数数组，数组里面可能含有相同的元素，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前面部分，所有
的偶数位于数组的后面部分，对奇数和奇数，偶数和偶数之间的相对位置不做要求，但是时间复杂度和空间复杂度必须如下要求。

要求：时间复杂度O(n)，空间复杂度O(1)
*/

func reOrderArrayTwo(array []int) []int {
	n := len(array)
	left, right := 0, 0
	for right < n {
		if array[right]%2 == 1 {
			right++
			left++
		} else {
			//right往后找到第一个奇数，然后和left交换
			for right < n && array[right]%2 == 0 {
				right++
			}
			if right < n {
				array[right], array[left] = array[left], array[right]
				left++
			}
		}
	}
	return array
}
