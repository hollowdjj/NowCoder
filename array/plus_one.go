package array

/*
给定一个用数组表示的数字，即数组中每个数表示一个数位上的数，例如 [1,2,3]，表示 123。
请问给这个数字加一后得到的结果（结果同样以数组的形式返回）。
注意：数组中不可能出现负数，且保证数组的首位即数字的首位不可能是 0 。
*/

func PlusOne(nums []int) []int {
	n := len(nums)
	res := make([]int, n+1)
	carry := 1 //进位
	for i := n - 1; i >= 0; i-- {
		res[i+1] = (nums[i] + carry) % 10
		carry = (nums[i] + carry) / 10
	}
	if carry != 0 {
		res[0] = carry
		return res
	}
	return res[1:]
}
