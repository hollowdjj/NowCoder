package codetop

/*
给你一个nxn的二进制矩阵grid中，返回矩阵中最短畅通路径的长度。如果不存在这样的路径，返回-1 。

二进制矩阵中的畅通路径是一条从左上角单元格（即，(0, 0)）到右下角单元格（即，(n - 1, n - 1)）的路径，该路径
同时满足下述要求：

路径途经的所有单元格都的值都是 0 。
路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
畅通路径的长度 是该路径途经的单元格总数。
*/

func shortestPathBinaryMatrix(grid [][]int) int {
	/*
		此题就是一个简单的单源最短路径问题。BFS和DFS都可以求解，但是最好采用BFS求解，DFS可能会超时。
		BFS与DFS的适用情景分析：
		1.如果只是要找到某一个结果是否存在，那么DFS会更高效。因为DFS会首先把一种可能的情况尝试到底，才
		会回溯去尝试下一种情况，只要找到一种情况，就可以返回了。但是BFS必须所有可能的情况同时尝试，在找
		到一种满足条件的结果的同时，也尝试了很多不必要的路径；
		2.如果是要找所有可能结果中最短的，那么BFS回更高效。因为DFS是一种一种的尝试，在把所有可能情况尝
		试完之前，无法确定哪个是最短，所以DFS必须把所有情况都找一遍，才能确定最终答案（DFS的优化就是剪
		枝，不剪枝很容易超时）。而BFS从一开始就是尝试所有情况，所以只要找到第一个达到的那个点，那就是
		最短的路径，可以直接返回了，其他情况都可以省略了，所以这种情况下，BFS更高效。
	*/
	if grid[0][0] == 1 {
		return -1
	}
	n := len(grid)
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	//BFS
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	step := 0
	q := [][2]int{{0, 0}}
	visited[0][0] = true
	for len(q) > 0 {
		step++
		size := len(q)
		for i := 0; i < size; i++ {
			front := q[0]
			q = q[1:]
			if front[0] == n-1 && front[1] == n-1 {
				return step
			}
			for i := range dirs {
				ni, nj := front[0]+dirs[i][0], front[1]+dirs[i][1]
				if ni < 0 || ni >= n || nj < 0 || nj >= n || grid[ni][nj] == 1 || visited[ni][nj] {
					continue
				}
				visited[ni][nj] = true
				q = append(q, [2]int{ni, nj})
			}
		}
	}
	return -1
}
