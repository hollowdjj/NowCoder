package jz_II

/*
给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。
如果一个不存在，返回0 。

输入: arr = [1,2,3,4,5,6,7,8]
输出: 5
解释: 最长的斐波那契式子序列为 [1,2,3,5,8]
*/

func lenLongestFibSubseq(arr []int) int {
	//dp[i][j]表示以arr[i],arr[j]结尾的最长斐波那契子序列长度
	//状态转移方程为：
	//dp[i][j] = Max{dp[k][i], where k<i and arr[k]+arr[i]=arr[j]}
	//使用哈希表保证arr[k]，可将时间复杂度降至O(n^2)
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dic := map[int]int{arr[0]: 0}
	res := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if k, ok := dic[arr[j]-arr[i]]; ok {
				dp[i][j] = max(dp[k][i], 2) + 1
				res = max(res, dp[i][j])
			}
		}
		dic[arr[i]] = i
	}
	return res
}
