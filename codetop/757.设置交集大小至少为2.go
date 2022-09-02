package codetop

import "sort"

/*
一个整数区间 [a, b]  ( a < b ) 代表着从 a 到 b 的所有连续整数，包括 a 和 b。
给你一组整数区间intervals，请找到一个最小的集合 S，使得 S 里的元素与区间intervals
中的每一个整数区间都至少有2个元素相交。输出这个最小集合S的大小。

输入: intervals = [[1, 3], [1, 4], [2, 5], [3, 5]]
输出: 3
解释:
考虑集合 S = {2, 3, 4}. S与intervals中的四个区间都有至少2个相交的元素。
且这是S最小的情况，故我们输出3。
*/
func intersectionSizeTwo(intervals [][]int) int {
	//按照intervals[i][0]升序排列，相同时按照intervals[i][1]降序排列
	//这样排序的使得起点相同的区间，大区间在前面，从而当后面的小区间满足
	//交集有至少2个元素时，大区间也一定满足
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	//从后往前遍历数组，初始时取区间[intervals[n-1][0],intervals[n-1][0]+1]
	//与区间intervals[n-1]的相交集合刚好有两个元素。
	n := len(intervals)
	start, end := intervals[n-1][0], intervals[n-1][0]+1
	res := 2
	for i := len(intervals) - 2; i >= 0; i-- {
		x, y := intervals[i][0], intervals[i][1]
		//有三种情况：
		//1. y >= end。说明intervals[i]是个大区间，所以一定满足相交集合至少有2个元素，集合大小不变
		//2. y < start。此时区间没有交集，贪心取start=x, end=x+1，集合大小+2
		//3. start <= y < end。此时只有一个交集，贪心取end=start,start=x，集合大小+1
		if y >= end {
			continue
		}
		if y < start {
			start, end = x, x+1
			res += 2
		} else if start <= y && y < end {
			start, end = x, start
			res++
		}
	}
	return res
}
