package codetop

/*
给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。

请你返回让 s 成为回文串的 最少操作次数 。

「回文串」是正读和反读都相同的字符串。

输入：s = "mbadm"
输出：2
解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
*/

func minInsertions(s string) int {
	//dp[i][j]表示使得s[i...j]成为回文串的最少插入次数，状态转移方程为：
	//1. s[i] == s[j]时，不需要插入 dp[i][j] = dp[i+1][j-1]
	//2. s[i] != s[j]时，我们需要判断是插入s[i]还是s[j]使得插入次数最少，即dp[i][j]=min(dp[i+1][j],dp[i][j-1])+1
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}
