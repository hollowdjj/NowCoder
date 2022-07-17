package jz_II

/*
现在总共有 numCourses 门课需要选，记为 0 到 numCourses-1。
给定一个数组 prerequisites ，它的每一个元素 prerequisites[i] 表示两门课程之间的先修顺序。
例如 prerequisites[i] = [ai, bi] 表示想要学习课程 ai ，需要先完成课程 bi 。
请根据给出的总课程数  numCourses 和表示先修顺序的 prerequisites 得出一个可行的修课序列。
可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

输入: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该
排在课程 0 之后。因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	//视为一个图。从而问题变为能否找到一个起点，使得从其开始的路径能够
	//串联起图中所有的节点，也即一个拓扑排序。

	//统计所有节点的入度
	indeg := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i, _ := range prerequisites {
		start, end := prerequisites[i][1], prerequisites[i][0]
		graph[start] = append(graph[start], end)
		indeg[end]++
	}

	//将所有入度为0的节点放入队列中（这部分节点的顺序无关紧要）
	var q []int
	for i, n := range indeg {
		if n == 0 {
			q = append(q, i)
		}
	}
	//BFS
	var res []int
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		res = append(res, front)
		for _, node := range graph[front] {
			indeg[node]--
			if indeg[node] == 0 {
				q = append(q, node)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return []int{}
}
