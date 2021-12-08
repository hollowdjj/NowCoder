package array

/*
矩阵的最小路径和
给定一个n*m的矩阵a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是
路径和，输出所有路径中最小的路径和。
要求：空间复杂度O(1)，时间复杂度O(nm)
*/

//MinPathSum 动态规划解决矩阵的最小路径和问题
func MinPathSum(matrix [][]int) int {
	//dp数组的定义为：dp[i][j]为到matrix[i][j]的最小路径和
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	//base case
	dp[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	//选择
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = matrix[i][j] + MinOfTwoInt(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[n-1][m-1]
}
