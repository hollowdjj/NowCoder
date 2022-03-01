package array

/*
连续子数组的最大乘积

输入一个长度为n的整型数组nums，数组中的一个或连续多个整数组成一个子数组。求所有子数组的乘积的最大值。
1.子数组是连续的，且最小长度为1，最大长度为n
2.长度为1的子数组，乘积视为其本身，比如[4]的乘积为4

输入：[-3,2,-2]
输出：[12]
*/

func MaxProduct(nums []int) int {
	//dp[i][0]和dp[i][1]表示以nums[i]结尾的连续子串的最大乘积和最小乘积，那么状态转移方程为：
	//dp[i][0] = Max(nums[i],Max(nums[i]*dp[i-1][0],nums[i]*dp[i-1][1]))
	//dp[i][1] = Min(nums[i],Min(nums[i]*dp[i-1][0],nums[i]*dp[i-1][1]))
	//由于dp[i]只和dp[i-1]有关，所有可以进行状态压缩，使得空间复杂度变为O(1)
	n := len(nums)
	minVal, maxVal := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		maxVal, minVal = max(nums[i], max(nums[i]*minVal, nums[i]*maxVal)), min(nums[i], min(nums[i]*minVal, nums[i]*maxVal))
		res = max(res, maxVal)
	}
	return res
}
