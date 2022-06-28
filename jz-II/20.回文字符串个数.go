package jz_II

/*
给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
*/

func countSubstrings(s string) int {
	//dp[i][j]表示s[i..j]是否是回文子串。状态转移方程为：
	//1. s[i] != s[j]  dp[i][j]=false
	//2. s[i] == s[j]  dp[i][j]=dp[i+1][j-1]
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		//dp[i][i]=true
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}
