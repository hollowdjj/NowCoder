package others

/*
01背包问题
已知一个背包最多能容纳体积之和为v的物品
现有 n 个物品，第 i 个物品的体积为 vi , 重量为 wi
求当前背包最多能装多大重量的物品?

输入：10,2,[[1,3],[10,4]]
返回值：4
说明：第一个物品的体积为1，重量为3，第二个物品的体积为10，重量为4。只取第二个物品可以达到最优方案，取物重量为4
*/

func Knapsack(V int, n int, vw [][]int) int {
	//dp[i][j]表示只有i个物品时，背包容量为j时能装的最大重量
	//dp[i][j] = Max{dp[i-1][j],dp[i][dp[j-vw[i][0]]]+vw[i][1]}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, V+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= V; j++ {
			if t := j - vw[i-1][0]; t >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][t]+vw[i-1][1])
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}

	return dp[n][V]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
