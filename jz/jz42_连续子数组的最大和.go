package jz

/*
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，子数组最小长度为1。求所有子数组的和的最大值。

输入：[1,-2,3,10,-4,7,2,-5]
输出：18
说明：经分析可知，输入数组的子数组[3,10,-4,7,2]可以求得最大和为18
*/

func FindGreatestSumOfSubArray(array []int) int {
	//dp[i]表示以array[i]结尾的子数组的最大和，状态转移方程为：dp[i] = Max{array[i],dp[i-1]+array[i]}。
	//由于dp[i]只和dp[i-1]有关，因此可以状态压缩，使得空间复杂度降为O(1)
	n := len(array)
	prev := array[0]
	res := prev
	for i := 1; i < n; i++ {
		curr := max(array[i], prev+array[i])
		res = max(res, curr)
		prev = curr
	}
	return res
}
