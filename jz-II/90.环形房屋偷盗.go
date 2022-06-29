package jz_II

/*
一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都
围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组nums，请计算在不触动警报装置的情况下，今晚能够偷窃到的
最高金额。

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
*/
func rob1(nums []int) int {
	//首尾两间房不能同时偷，因此，分两种情况：1.不偷最后一间 2.不偷第一间；并取两种情况的最大值
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	//不偷最后一间房。
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n-1; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	res := dp[n-2]
	//不偷第一间房
	dp[1] = nums[1]
	dp[2] = max(nums[1], nums[2])
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	if dp[n-1] > res {
		res = dp[n-1]
	}
	return res
}
