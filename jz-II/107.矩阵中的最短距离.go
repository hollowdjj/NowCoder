package jz_II

/*
给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置
元素到最近的 0 的距离。两个相邻元素间的距离为 1 。
*/

func updateMatrix(mat [][]int) [][]int {
	//使用一个"超级零"，这个超级零到矩阵中所有0的距离为0。从而，求所有
	//1到这个超级零的最短距离，就是所有1到最近的0的距离，进而变成了一个单源
	//最短路径问题。
	n, m := len(mat), len(mat[0])
	dis := make([][]int, n)
	visit := make([][]int, n)
	q := make([][2]int, 0)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, m)
		visit[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				visit[i][j] = 1
				q = append(q, [2]int{i, j})
			}
		}
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			i, j := front[0]+dirs[i][0], front[1]+dirs[i][1]
			if i < 0 || i >= n || j < 0 || j >= m || visit[i][j] == 1 {
				continue
			}
			dis[i][j] = dis[front[0]][front[1]] + 1
			visit[i][j] = 1
			q = append(q, [2]int{i, j})
		}
	}
	return dis
}
