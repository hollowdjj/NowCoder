package array

/*
岛屿的数量
给一个01矩阵，1代表是陆地，0代表海洋， 如果两个1相邻，那么这两个1属于同一个岛。我们只考虑上下左右为相邻。
岛屿: 相邻陆地可以组成一个岛屿（相邻:上下左右） 判断岛屿个数。
例如：
输入
[
[1,1,0,0,0],
[0,1,0,1,1],
[0,0,0,1,1],
[0,0,0,0,0],
[0,0,1,1,1]
]
对应的输出为3
*/

func IslandNum(grid [][]byte) int {
	//遍历数组，当遇到1的时候，岛屿数量加1，并使用深度优先搜索将它上下左右的1，以及这些1
	//上下左右的1全部置为0
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfsIsland(grid, i, j)
			}
		}
	}
	return res
}

func dfsIsland(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfsIsland(grid, i-1, j) //上
	dfsIsland(grid, i+1, j) //下
	dfsIsland(grid, i, j-1) //左
	dfsIsland(grid, i, j+1) //右
}
