package jz_II

/*
一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响小偷偷窃的唯一制约因素就是相邻的
房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组 nums，请计算不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

输入：nums = [1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。偷窃到的最高金额 = 1 + 3 = 4 。
*/
func rob(nums []int) int {
	//dp[i]表示只在前i个房屋偷盗，能盗取的最高金额，状态转移方程为：
	//dp[i] = Max{dp[j], where 0<=j<i-1} + nums[i]
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		prevMax := 0
		for j := 0; j < i-1; j++ {
			if dp[j] > prevMax {
				prevMax = dp[j]
			}
		}
		dp[i] = max(prevMax+nums[i], dp[i-1])
	}
	return dp[n-1]
}
