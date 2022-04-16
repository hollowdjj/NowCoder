package jz

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长的m段（m和n都是整数，n > 1并且m > 1，m <= n ）
每段绳子的长度记为 k[1],...,k[m] 。请问 k[1]*k[2]*...*k[m] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为 2、3、3 的三段，此时得到的最大乘积是18。
*/

func cutRope(n int) int {
	//dp[i]表示当绳子长度为i时的最大乘积。假设将绳子分成j和i-j两段，那么有四种情况
	//1. j和i-j都不再细分，dp[i] = j*(i-j)
	//2. j细分，但i-j不细分，dp[i] = dp[j]*(i-j)
	//3. j不细分，但i-j细分，dp[i] = j*dp[i-j]
	//4. j和i-1都细分，dp[i] = dp[j] * dp[i-j]
	//显然，dp[i]是取上面四个值中的最大值。因此，状态转移方程为：
	//dp[i] = max(dp[i],max(j,dp[j]) * max(dp[i-j],i-j))
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j, dp[j])*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
