package jz_II

/*
树可以看成是一个连通且无环的无向图。给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个
顶点包含在1到n中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为n的二维数组edge，dges[i]=[ai,bi]
表示图中在ai和bi之间存在一条边。请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多
个答案，则返回数组edges中最后出现的边。

输入: edges = [[1,2],[1,3],[2,3]]
输出: [2,3]
*/

func findRedundantConnection(edges [][]int) []int {
	//使用并查集求解。具体方法为：开始时，每个节点为一个集合。然后遍历每一条边，然后合并集合
	//如果合并时，发现两个点已经在同一个集合中了，那么这条边就是多余的边。
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if x == parent[x] {
			return parent[x]
		}
		parent[x] = find(parent[x])
		return parent[x]
	}
	merge := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for i := range edges {
		if !merge(edges[i][0], edges[i][1]) {
			return edges[i]
		}
	}
	return nil
}
