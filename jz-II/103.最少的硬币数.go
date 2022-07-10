package jz_II

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少
的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。你可以认为每种硬币的数量是无限的。

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
*/

func coinChange(coins []int, target int) int {
	//这是一个完全背包问题，因为每个硬币可以选无数次。0-1背包问题与
	//完全背包问题的状态转移方程很相似:
	//1. 0-1背包：dp[i][j] = min(dp[i-1][j],dp[i-1][j-coins[i-1]]+1)
	//2. 完全背包：dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
	//一个是dp[i-1][j-coins[i-1]]，另一个是dp[i][j-coins[i-1]]。这是因为
	//在0-1背包中，每个元素只能使用一次，所以是i-1，而完全背包可以使用无数次，所以是i
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = target + 1
		}
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j >= coins[i-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	res := dp[n][target]
	if res == target+1 {
		res = -1
	}
	return res
}

func coinChange1(coins []int, target int) int {
	//可以看到dp[i]只与dp[i-1]有关，因此可以使用一维滚动数组进行状态压缩
	//此时，dp[i]可以理解为，使用所有硬币凑i的最少硬币数
	dp := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = target + 1
	}
	dp[0] = 0
	n := len(coins)
	for i := 1; i <= target; i++ {
		for j := 0; j < n; j++ {
			if i >= coins[j] {
				//这里dp[i]的当前值其实就是dp[i-1][j]
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[target] == target+1 {
		return -1
	}
	return dp[target]
}
