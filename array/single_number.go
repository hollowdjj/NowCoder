package array

/*
给定一个整数数组，数组中有一个数出现了一次，其他数出现了两次，请找出只出现了一次的数。
*/

func SingleNumber(nums []int) int {
	//此题需要根据异或(相同为0，不同为1)的性质来做：
	//1. 任何数异或上0都不会改变
	//2. 异或运算满足交换律，即a⊕b⊕a = (a⊕a)⊕b  = 0⊕b = b
	//3. 两个相同的数异或为0，即a⊕a = 0
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
