package jz_II

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4]
*/

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(path []int, count int, i int)
	dfs = func(path []int, count int, i int) {
		if count == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(path, count+1, j+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 1)
	return res
}
