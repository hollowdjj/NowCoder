package jz

/*
给一个长度为 n 的数组，数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。
*/

func MoreThanHalfNum(numbers []int) int {
	//黑帮火并算法
	candidates, count := 0, 0
	n := len(numbers)
	for i := 0; i < n; i++ {
		if count == 0 {
			candidates = numbers[i]
			count++
		} else {
			if candidates == numbers[i] {
				count++
			} else {
				count--
			}
		}
	}
	return candidates
}
