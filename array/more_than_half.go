package array

/*
给一个长度为 n 的数组，数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。
要求：空间复杂度：O(1)，时间复杂度O(n)
*/

func MoreThanHalfNum(numbers []int) int {
	/*
		候选人法。当两个数字不想等时，同时消除这两个数字，最坏的情况下是同时消除一个众数和非众数。
		因此，最后留下来的数字一定就是众数。但，需要注意的是，我们不能真的去消除数组中的数字，只能
		想办法用其他方法代替。
	*/
	candidate, count := 0, 0
	for i := 0; i < len(numbers); i++ {
		//当candidate候选人的数量为0时，选择当前元素作为候选人
		if count == 0 {
			candidate = numbers[i]
		}
		if numbers[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
