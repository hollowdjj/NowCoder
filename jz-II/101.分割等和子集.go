package jz_II

/*
给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。

输入：nums = [1,5,11,5]
输出：true
解释：nums 可以分割成 [1, 5, 5] 和 [11] 。

输入：nums = [1,2,3,5]
输出：false
解释：nums 不可以分为和相等的两部分
*/

func canPartition(nums []int) bool {
	//题目可以转换为，在nums中找若干个数使其和为nums中所有数字之和sum的二分之一。
	//因而是一个01背包问题，使用动态规划求解。dp[i][j]表示从nums中前i个数字中选取
	//数字，是否能凑出j。
	sum := 0
	for _, r := range nums {
		sum += r
	}
	if sum%2 == 1 {
		return false
	}
	n := len(nums)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum/2+1)
		dp[i][0] = true //base case。无论如何都可以凑出0
	}
	//选择
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum/2; j++ {
			//选择nums[i-1]
			dp[i][j] = dp[i-1][j]
			if !dp[i][j] && j >= nums[i-1] {
				dp[i][j] = dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum/2]
}
