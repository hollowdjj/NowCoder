package array

/*
牛妹给了牛牛一个长度为n的下标从0开始的正整型数组a ，粗心的牛牛不小心把其中的一些数字删除了。
假如a_i被删除了,则a_i=0。对于所有被删除的数字，牛牛必须选择一个正整数填充上。现在牛牛想知道有多少种填充方案使得：
a0≤a1≤a2...≤an-1，且1≤ai≤k
函数传入一个下标从0开始的数组a和一个正整数k ，请返回合法的填充方案数对 10^9+7取模的值,保证不存在方案数为0的数据。

输入：[0,4,5],6
返回值：4
说明：所有的合法填充方案是：[1,4,5],[2,4,5],[3,4,5],[4,4,5]，共4种。
*/

func FillArray(a []int, k int) int {
	length := len(a)
	res := 0

	//left为0序列前一个位置，right为0序列后一个位置
	left, right := 0, 0
	for right < length {
		if a[right] != 0 {
			//得到可填充数字的上下限
			lb, ub := 1, k
			if left >= 1 {
				lb = a[left-1]
			}
			if right < length-1 {
				ub = a[right]
			}
			//一共要填充right - left个0。ub到lb之间，一共有ub-lb+1个数
			m := right - left
			n := ub - lb + 1
			temp := factorial(n) / factorial(n-m)
			res *= temp
			left = right + 1
		}
		right++
	}

	return res % 1000000007
}

func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}
