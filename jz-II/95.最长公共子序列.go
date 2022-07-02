package jz_II

/*
给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列的长度。如果不存在公共子序列，返回0。

一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符
（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
*/

func longestCommonSubsequence(str1 string, str2 string) int {
	//dp[i][j]表示str1[0:i]与str2[0:j]的最长公共子序列长度
	//状态转移方程为：
	//1. str1[i-1] == str2[j-1], dp[i][j] = dp[i-1][j-1] + 1
	//2. str1[i-1] != str2[j-1], dp[i][j] = max(dp[i-1][j],dp[i][j-1])
	n1, n2 := len(str1), len(str2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	res := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			res = max(res, dp[i][j])
		}
	}

	return res
}
