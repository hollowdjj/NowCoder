package jz_II

import "math"

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是下标与上一层结点下标相同或者
等于上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移
动到下一行的下标 i 或 i + 1 。

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
*/

func minimumTotal(triangle [][]int) int {
	//dp[i][j]表示走到triangle[i][j]的最小路径
	n := len(triangle)
	dp := make([][]int, n)
	dp[0] = []int{triangle[0][0]}
	for i := 1; i < n; i++ {
		m := len(triangle[i])
		dp[i] = make([]int, m)
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < m-1; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][m-1] = dp[i-1][m-2] + triangle[i][m-1]
	}
	//找出最后一行的最小值
	res := math.MaxInt32
	m := len(triangle[n-1])
	for j := 0; j < m; j++ {
		if res > dp[n-1][j] {
			res = dp[n-1][j]
		}
	}
	return res
}
