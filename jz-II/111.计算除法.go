package jz_II

func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	//视作一个无向图，a/b视为a --- b, a/b的结果就是边的权重。
	//对每一个query，其实就是找到query[0]到query[1]的路径长度，因此需要求每个点
	//到其余点的最短路径，是一个多源最短路径问题，考虑使用Floyd算法。

	//对equations中的每个元素编号
	id := make(map[string]int)
	for _, e := range equations {
		e1, e2 := e[0], e[1]
		if _, hit := id[e1]; !hit {
			id[e1] = len(id)
		}
		if _, hit := id[e2]; !hit {
			id[e2] = len(id)
		}
	}

	//初始化邻接矩阵
	n := len(id)
	grid := make([][]float64, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]float64, n)
		grid[i][i] = 1
	}

	for i, e := range equations {
		e1, e2 := e[0], e[1]
		grid[id[e1]][id[e2]] = values[i]
		grid[id[e2]][id[e1]] = 1.0 / values[i]
	}

	//floyd算法
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			//计算grid[i][j]的方式是通过依次枚举中继点的方式
			for j := 0; j < n; j++ {
				if grid[i][k] > 0 && grid[k][j] > 0 {
					grid[i][j] = grid[i][k] * grid[k][j]
				}
			}
		}
	}

	var res []float64
	for _, q := range queries {
		q1, q2 := q[0], q[1]
		_, hit1 := id[q1]
		_, hit2 := id[q2]
		if !hit1 || !hit2 || grid[id[q1]][id[q2]] == 0 {
			res = append(res, -1)
		} else {
			res = append(res, grid[id[q1]][id[q2]])
		}
	}

	return res
}

func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	//使用邻接链表+BFS

	//对equations中的每个元素编号
	id := make(map[string]int)
	for _, e := range equations {
		e1, e2 := e[0], e[1]
		if _, hit := id[e1]; !hit {
			id[e1] = len(id)
		}
		if _, hit := id[e2]; !hit {
			id[e2] = len(id)
		}
	}

	//构建邻接链表
	type edge struct {
		to     int
		weight float64
	}
	graph := make([][]edge, len(id))
	for i, e := range equations {
		start, end := id[e[0]], id[e[1]]
		graph[start] = append(graph[start], edge{end, values[i]})
		graph[end] = append(graph[end], edge{start, 1 / values[i]})
	}
	//BFS
	bfs := func(start, end int) float64 {
		q := []int{start}
		path := make([]float64, len(graph)) //path[i]表示start到i的路径总权重
		path[start] = 1
		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			if front == end {
				return path[front]
			}
			//遍历邻接链表
			for _, e := range graph[front] {
				//避免走回头路
				if path[e.to] == 0 {
					path[e.to] = path[front] * e.weight
					q = append(q, e.to)
				}
			}
		}
		return -1
	}

	//对每一个query执行BFS
	res := make([]float64, len(queries))
	for i, query := range queries {
		start, hit1 := id[query[0]]
		end, hit2 := id[query[1]]
		if !hit1 || !hit2 {
			res[i] = -1
		} else {
			res[i] = bfs(start, end)
		}
	}
	return res
}
