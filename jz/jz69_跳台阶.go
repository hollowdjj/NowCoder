package jz

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
*/

func jumpFloor(number int) int {
	//dp[i]表示调到第i级台阶的跳法，故状态转移方程为：
	//dp[i] = dp[i-1] + dp[i-2]
	dp := make([]int, number+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}
