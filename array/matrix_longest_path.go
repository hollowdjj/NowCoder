package array

/*
矩阵最长递增路径

给定一个 n 行 m 列矩阵 matrix ，矩阵内所有数均为非负整数。 你需要在矩阵中找到一条最长路径，使这条路径上的元素是递增的。
并输出这条最长路径的长度。
这个路径必须满足以下条件：
1. 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外。
2. 你不能走重复的单元格。即每个格子最多只能走一次。
数据范围：1≤n,m≤1000，0≤matrix[i][j]≤1000

例如矩阵：
1  2  3
4  5  6
7  8  9
返回值：5
说明：1->2->3->6->9即可。当然这种递增路径不是唯一的。
*/

func MatrixLongestPath(matrix [][]int) int {
	//使用深度优先算法
	row, col := len(matrix), len(matrix[0])
	var dfs func(matrix [][]int, i, j int, prev int) int
	dfs = func(matrix [][]int, i, j int, prev int) int {
		if i < 0 || i >= row || j < 0 || j >= col || matrix[i][j] <= prev {
			//这里的matrix[i][j]<=prev的判断条件就对应了题目中的路径元素必须是递增的
			return 0
		}
		ret := 0
		ret = max(ret, dfs(matrix, i+1, j, matrix[i][j])) //向上
		ret = max(ret, dfs(matrix, i-1, j, matrix[i][j])) //向下
		ret = max(ret, dfs(matrix, i, j-1, matrix[i][j])) //向左
		ret = max(ret, dfs(matrix, i, j+1, matrix[i][j])) //向右

		return ret + 1 //路径长度要算上matrix[i][j]这个元素
	}

	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res = max(res, dfs(matrix, i, j, -1))
		}
	}
	return res
}

func MatrixLongestPathAdvance(matrix [][]int) int {
	//dfs+dp。我们注意到在上面的深度优先中可以存储计算过的值，从而加速搜索。
	//dp[i][j]表示以matrix[i][j]为头的最长递增路径的长度
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	var dfs func(matrix [][]int, i, j int, prev int) int
	dfs = func(matrix [][]int, i, j int, prev int) int {
		if i < 0 || i >= row || j < 0 || j >= col || matrix[i][j] <= prev {
			//这里的matrix[i][j]<=prev的判断条件就对应了题目中的路径元素必须是递增的
			return 0
		}
		if dp[i][j] > 0 {
			//当前值已经计算过了，直接返回
			return dp[i][j]
		}
		ret := 0
		ret = max(ret, dfs(matrix, i+1, j, matrix[i][j])) //向上
		ret = max(ret, dfs(matrix, i-1, j, matrix[i][j])) //向下
		ret = max(ret, dfs(matrix, i, j-1, matrix[i][j])) //向左
		ret = max(ret, dfs(matrix, i, j+1, matrix[i][j])) //向右

		dp[i][j] = ret + 1
		return ret + 1 //路径长度要算上matrix[i][j]这个元素
	}
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res = max(res, dfs(matrix, i, j, -1))
		}
	}
	return res
}
