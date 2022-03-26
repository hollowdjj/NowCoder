package array

/*
给定一个整数数组 nums 表示不同数额的硬币和一个正整数 target 表示总金额，请你计算
并返回可以凑出总金额的的组合数。如果凑不出 target 则返回 0。假设每一种面额的硬币有
无数多个

输入：  5,[1,2,4,5]
返回值：5
*/

//ChangeDfs DFS解法，时间复杂度为O(n!)
func ChangeDfs(target int, nums []int) int {
	res := 0
	var dfs func(i int, rest int)
	dfs = func(i int, rest int) {
		if rest < 0 {
			return
		} else if rest == 0 {
			res += 1
		}
		for j := i; j < len(nums); j++ {
			dfs(j, rest-nums[j])
		}
	}

	return res
}

func ChangeDp(target int, nums []int) int {
	//dp[i][j]表示用第i及其之前所有的硬币凑j，能有几种凑法。从而，状态转移方程为：
	//			dp[i][j] = dp[i-1][j] + dp[i][j-nums[i]]
	//dp[i-1][j]		表示的是不用第i枚硬币。那么需要凑的数不变，还是j
	//dp[i][j-nums[i]]	表示使用第i枚硬币。那么此时现在需要凑的数变成了j-nums[i]
	//显然，对于一枚硬币，只有选还是不选两种可能，故状态转移方程就是以上两项相加
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		//这里从j从nums[i]开始的原因是，当剩余金额小于当前硬币面额时，现有组合数不会发生变化
		for j := nums[i]; j <= target; j++ {
			if i-1 >= 0 {
				dp[i][j] += dp[i-1][j] + dp[i][j-nums[i]]
			} else {
				dp[i][j] += dp[i][j-nums[i]]
			}
		}
	}

	return dp[n-1][target]
}

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。
你可以认为每种硬币的数量是无限的。
*/

func CoinChange(coins []int, amount int) int {
	n := len(coins)

	//dp[i]表示使用coins凑出金额i所需的最少硬币数。那么状态转移方程为：
	//dp[i] = Min{dp[i-coins[j]]}，其中0<=j<len(coins)-1
	//dp[0] = 0
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if i < coins[j] {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
