package jz_II

import "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func merge(intervals [][]int) [][]int {
	//按照左端点大小的进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	n := len(intervals)
	for i := 1; i < n; i++ {
		last := res[len(res)-1]
		curr := intervals[i]
		if curr[0] > last[1] {
			res = append(res, curr)
		} else if curr[0] <= last[1] && curr[1] > last[1] {
			//形如[2,4][3,5]
			res[len(res)-1] = []int{last[0], curr[1]}
		}
	}
	return res
}
