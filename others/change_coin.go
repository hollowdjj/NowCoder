package others

/*
兑换零钱(一)
给定数组arr，arr中所有的值都为正整数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张，再
给定一个aim，代表要找的钱数，求组成aim的最少货币数。如果无解，请返回-1.

数据范围：数组大小满足0≤n≤10000，数组中每个数字都满足0<val≤10000，0≤aim≤5000
*/

func ChangeCoin(arr []int, aim int) int {
	//dp[i]表示当目标金额为i时，组成i的最少货币数。故，状态转移方程为：
	//dp[i] = Min{dp[i-arr[j]]}，其中j = 0....len(arr)-1
	dp := make([]int, aim+1)
	for i := 0; i <= aim; i++ {
		dp[i] = aim + 1
	}
	dp[0] = 0

	for i := 0; i <= aim; i++ {
		for _, v := range arr {
			if i < v {
				continue
			}
			dp[i] = min(dp[i], dp[i-v])
		}
	}

	if dp[aim] == aim+1 {
		return -1
	}
	return dp[aim]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
