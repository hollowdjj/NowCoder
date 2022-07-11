package jz_II

/*
给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
一个岛屿是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个1必须在水平或者竖直方向
上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。找到给定的二维数组中最大的岛屿面积。
如果没有岛屿，则返回面积为0。
*/

func maxAreaOfIsland(grid [][]int) int {
	//注意，此题是求最大岛屿面积，而非岛屿个数
	n, m := len(grid), len(grid[0])
	res, curr := 0, 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 0 {
			return
		}
		curr++
		grid[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
			res = max(res, curr)
			curr = 0
		}
	}
	return res
}
