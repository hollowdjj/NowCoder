package jz_II

/*
给定一个由不同正整数组成的数组nums，和一个目标整数target。请从nums中找出并返回总和为target
的元素组合的个数。数组中的数字可以在一次排列中出现任意次，但是顺序不同的序列被视作不同的组合。
题目数据保证答案符合32位整数范围。

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
*/

func combinationSum4(nums []int, target int) int {
	//dp[i][j]表示在前i个数中选，凑出j的方案数，状态转移方程为：
	//dp[i][j] = dp[i-1][j] + dp[i][j-nums[i-1]]
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		for j := 1; j <= target; j++ {
			//在前i个数中和为j-nums[k-1]的序列加上nums[k-1]
			//可得和j
			for k := 1; k <= i; k++ {
				if j >= nums[k-1] {
					dp[i][j] += dp[i][j-nums[k-1]]
				}
			}
		}
	}

	return dp[n][target]
}

func combinationSum4Advance(nums []int, target int) int {
	//状态压缩
	n := len(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < n; j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
