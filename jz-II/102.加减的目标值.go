package jz_II

/*
给定一个正整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表
达式 "+2-1" 。返回可以通过上述方法构造的、运算结果等于 target 的不同表达式的数目。
*/

func findTargetSumWays(nums []int, target int) int {
	//回溯法
	res, n := 0, len(nums)
	var backtrack func(sum int, i int)
	backtrack = func(sum int, i int) {
		if i == n {
			if sum == target {
				res++
			}
			return
		}
		//加
		backtrack(sum+nums[i], i+1)
		//减
		backtrack(sum-nums[i], i+1)
	}
	backtrack(0, 0)
	return res
}

func findTargetSumWays1(nums []int, target int) int {
	//本题的标注做法是视作0-1背包问题。设所有符号为"+"的数字的和为
	//positive，数组nums中所有元素的和为sum。那么所有符号为"-"的数字
	//的和为sum - positive = negative。positive - negative = target
	//进而有sum - 2 * negative = target, negative = (sum-target)/2。
	//因此，题目变为了：在nums中，找到任意n个数字，满足其和为negtive，问有多少
	//个这样的组合，即一个典型的0-1背包问题。
	sum := 0
	for _, v := range nums {
		sum += v
	}
	temp := sum - target
	if temp < 0 || temp%2 == 1 {
		return 0
	}
	neg, n := temp/2, len(nums)
	//dp[i][j]表示从前i个数字中选，凑出j的方案数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= neg; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][neg]
}
