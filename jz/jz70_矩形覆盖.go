package jz

/*
我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，从
同一个方向看总共有多少种不同的方法？

注意：约定 n == 0 时，输出 0
*/

func rectCover(number int) int {
	//dp[i]表示覆盖2*i的大矩形的覆盖方法，状态转移方程为：
	//dp[i] = dp[i-1] + dp[i-2]即横放和竖放两种
	if number == 0 {
		return 0
	}
	dp := make([]int, number+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}
