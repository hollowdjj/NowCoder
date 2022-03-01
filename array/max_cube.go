package array

/*
最大正方形
给定一个由'0'和'1'组成的2维矩阵，返回该矩阵中最大的由'1'组成的正方形的面积，输入的矩阵是字符形式而非数字形式。

数据范围：矩阵的长宽满足 0 \le n \le 200≤n≤20,矩阵中的元素属于 {'1','0'}
*/

func MaxCube(matrix [][]int) int {
	//dp[i][j]表示以matrix[i][j]为右下角的最大正方形的边长
	//1.当matrix[i][j]==0时,dp[i][j]=0
	//2.当matrix[i][j]==1时,dp[i][j]=Min{dp[i-1][j-1,dp[i-1][j],dp[i][j-1]]} + 1
	//1  1  0
	//1  1  1
	//1  1  1
	if len(matrix) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, col+1)
	}

	res := 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			res = max(res, dp[i][j])
		}
	}

	return res * res
}
