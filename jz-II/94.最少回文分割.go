package jz_II

import "math"

/*
给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。
返回符合要求的 最少分割次数。

输入：s = "aab"
输出：1
解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
*/

func minCut(s string) int {
	//dp[i]表示s[0..i]的最少分割次数。此时，我们枚举j，如果s[j+1...i]也是回文并且
	//dp[j]可行时，dp[i] = dp[j]+1

	//预处理
	n := len(s)
	dic := make([][]bool, n)
	for i := range dic {
		dic[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || dic[i+1][j-1]) {
				dic[i][j] = true
			}
		}
	}

	dp := make([]int, n)
	for i := range dp {
		//如果s[0:i+1]本身就是回文，那么就不需要分割了
		if dic[0][i] {
			continue
		}
		dp[i] = math.MaxInt32 - 1
		for j := 0; j < i; j++ {
			if dic[j+1][i] && dp[i] != -1 && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}
