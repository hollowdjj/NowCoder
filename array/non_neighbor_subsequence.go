package array

/*
不相邻最大子序列和

给你一个数组，其长度为 n  ，在其中选出一个子序列，子序列中任意两个数不能有相邻的下标（子序列可以为空）

本题中子序列指在数组中任意挑选若干个数组成的数组。

输入：4,[4,2,3,5]
输出：9
说明：[4,3] [4,5] [2,5]中[4 5]元素之和最大
*/

func NonNeighbotSubsequnce(n int, arr []int) int64 {
	//dp[i]表示array[0:i]的最优值。那么状态转移方程为：
	//dp[i]=Max{dp[i-1],dp[i-2]+arr[i-1]}。因为对arr[i-1]就只有选或者不选两种选择，不选arr[i-1]则为dp[i-1]。
	//选择arr[i-1]时，由于不能选相邻数，所以是dp[i-2]+arr[i-1]。
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = arr[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i-1])
	}
	if dp[n] < 0 {
		return 0
	}
	return int64(dp[n])
}
