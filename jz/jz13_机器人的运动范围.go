package jz

/*
地上有一个 rows 行和 cols 列的方格。坐标从 [0,0] 到 [rows-1,cols-1] 。一个机器人从坐标 [0,0] 的格子开始
移动，每一次只能向左，右，上，下四个方向移动一格，但是不能进入行坐标和列坐标的数位之和大于 threshold 的格子。
例如，当 threshold 为 18 时，机器人能够进入方格   [35,37] ，因为 3+5+3+7 = 18。但是，它不能进入方格 [35,38]
因为 3+5+3+8 = 19 。请问该机器人能够达到多少个格子？
*/

func movingCount(threshold int, rows int, cols int) int {
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	res := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || digitSum(i, j) > threshold || visited[i][j] {
			return
		}
		visited[i][j] = true
		res++
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	dfs(0, 0)
	return res
}

func digitSum(a, b int) int {
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	for b > 0 {
		sum += b % 10
		b /= 10
	}
	return sum
}
