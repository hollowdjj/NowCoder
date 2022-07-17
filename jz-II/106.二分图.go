package jz_II

/*
存在一个无向图 ，图中有n个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。
给定一个二维数组graph ，表示图，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。
形式上，对于 graph[u] 中的每个v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图
同时具有以下属性：

不存在自环（graph[u] 不包含 u）。
不存在平行边（graph[u] 不包含重复值）。
如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两
个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。

如果图是二分图，返回 true ；否则，返回 false 。
*/

/*
染色法。任意选取一个节点，将其染为红色，然后将所有与其相连的节点染成绿色。
遍历到某一节点时，如果节点已经被染色了，且颜色和将要给他染的颜色不一致，那么
说明不能构成二分图。
*/

func isBipartiteBFS(graph [][]int) bool {
	color := make([]int, len(graph))
	bfs := func(start int) bool {
		//BFS 染色
		q := []int{start}
		color[start] = 1
		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			for _, node := range graph[front] {
				if color[node] == 0 {
					color[node] = -1 * color[front]
					q = append(q, node)
				} else if color[node] == color[front] {
					return false
				}
			}
		}
		return true
	}
	//因为图可能不是连通的，所以需要尽可能遍历每一个点
	for i := 0; i < len(graph); i++ {
		if color[i] != 0 {
			continue
		}
		if !bfs(i) {
			return false
		}
	}
	return true
}

func isBipartiteDFS(graph [][]int) bool {
	color := make([]int, len(graph))
	res := true
	var dfs func(start int, c int)
	dfs = func(start int, c int) {
		color[start] = c
		for i := 0; i < len(graph[start]) && res; i++ {
			node := graph[start][i]
			if color[node] == color[start] {
				res = false
				return
			} else if color[node] == 0 {
				dfs(node, -1*color[start])
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		if color[i] == 0 {
			dfs(i, 1)
			if res == false {
				return false
			}
		}
	}
	return true
}
