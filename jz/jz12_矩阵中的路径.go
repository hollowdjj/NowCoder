package jz

/*
请设计一个函数，用来判断在一个n乘m的矩阵中是否存在一条包含某长度为len的字符串所有字符的路径。路径可以从矩阵中的任意一个格子开始，
每一步可以在矩阵中向左，向右，向上，向下移动一个格子。如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
例如:
							a  b  c  e
							s  f  c  s
							a  d  e  e
矩阵中包含一条字符串"bcced"的路径，但是矩阵中不包含"abcb"路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路
径不能再次进入该格子。
*/

func hasPath(matrix [][]byte, word string) bool {
	n, m := len(matrix), len(matrix[0])
	used := make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//以matrix[i][j]作为起点进行深搜
			if matrix[i][j] == word[0] {
				if dfs(matrix, word, used, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(matrix [][]byte, word string, used [][]bool, i, j int, index int) bool {
	//递归终止条件
	if index == len(word) {
		return true
	}
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || used[i][j] || matrix[i][j] != word[index] {
		return false
	}

	//回溯
	used[i][j] = true
	res := dfs(matrix, word, used, i-1, j, index+1) || dfs(matrix, word, used, i+1, j, index+1) ||
		dfs(matrix, word, used, i, j-1, index+1) || dfs(matrix, word, used, i, j+1, index+1)
	used[i][j] = false
	return res
}
