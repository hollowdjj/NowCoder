package codetop

/*
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

9	9	4
6	6	8
2	1	1
输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
输出：4
解释：最长递增路径为 [1, 2, 6, 9]。
*/

func longestIncreasingPath(matrix [][]int) int {
	//dfs+词典记忆化搜索
	n, m := len(matrix), len(matrix[0])
	dic := make([][]int, n)
	for i := range dic {
		dic[i] = make([]int, m)
	}
	var dfs func(i, j int, prev int) int
	dfs = func(i, j int, prev int) int {
		if i < 0 || i >= n || j < 0 || j >= m || prev >= matrix[i][j] {
			return 0
		}
		if dic[i][j] != 0 {
			return dic[i][j]
		}
		left := dfs(i, j-1, matrix[i][j])
		right := dfs(i, j+1, matrix[i][j])
		up := dfs(i-1, j, matrix[i][j])
		down := dfs(i+1, j, matrix[i][j])

		dic[i][j] = max(max(left, right), max(up, down)) + 1
		return dic[i][j]
	}
	res := 0
	for i := range matrix {
		for j := range matrix[i] {
			if dic[i][j] > 0 {
				continue
			}
			//这里的实参prev必须是最小值，不能是matrix(i,j)
			res = max(dfs(i, j, -1), res)
		}
	}
	return res
}
