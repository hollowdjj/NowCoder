package jz

/*
在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。
也不知道每个数字重复几次。请找出数组中任意一个重复的数字。 例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，
那么对应的输出是2或者3。存在不合法的输入的话输出-1
*/

func Duplicate(numbers []int) int {
	//原地哈希。也可以一次遍历完成，这里分了两次，使得思路更加清晰
	n := len(numbers)
	for i := 0; i < n; i++ {
		if numbers[i] != numbers[numbers[i]] {
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
			i--
		}
	}

	for i := 0; i < n; i++ {
		if numbers[i] != i {
			return numbers[i]
		}
	}

	return -1
}
