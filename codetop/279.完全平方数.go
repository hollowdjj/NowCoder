package codetop

import (
	"math"
)

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
*/

func numSquares(n int) int {
	//先找到1~n中的所有完全平方数
	var nums []int
	m := int(math.Sqrt(float64(n)))
	for i := 1; i <= m; i++ {
		if i*i > n {
			break
		}
		nums = append(nums, i*i)
	}
	//fmt.Println(nums)
	m = len(nums)
	//0-1背包问题，dp[i][j]表示前i个数凑j的最小数量
	//dp[i][j] = min{dp[i-1][j],dp[i][j-nums[i-1]]+1}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
		dp[i][0] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if j >= nums[i-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-nums[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
