package jz

/*
在一个m\times nm×n的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并
每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
如输入这样的一个二维数组，
[
[1,3,1],
[1,5,1],
[4,2,1]
]
那么路径 1→3→5→2→1 可以拿到最多价值的礼物，价值为12
*/

func GiftMaxValue(grid [][]int) int {
	//dp[i][j]表示走到grid[i-1][j-1]能拿到的礼物的最大价值，从而状态转移方程为：
	//dp[i][j] = Max{dp[i][j-1],dp[i-1][j]} + grid[i-1][j-1]
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	//选择
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + grid[i-1][j-1]
		}
	}
	return dp[n][m]
}
