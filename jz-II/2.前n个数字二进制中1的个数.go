package jz_II

/*
给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

输出：n=5
输出；[0,1,1,2,1,2]
解释:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101
*/

func countBits(n int) []int {
	//核心在于求二进制数中1的个数。count函数同样使用于负数
	count := func(num int) int {
		res := 0
		for num > 0 {
			num &= num - 1
			res++
		}
		return res
	}
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = count(i)
	}
	return res
}
