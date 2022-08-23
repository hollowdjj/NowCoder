package jz_II

/*
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符
串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
题目数据保证答案符合 32 位带符号整数范围。

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
*/

func numDistinct(s string, t string) int {
	//dp[i][j]表示s[:i]中t[:j]子序列的个数。那么状态转移方程为：
	//1. s[i-1] != t[j-1] dp[i][j] = dp[i-1][j]
	//2. s[i-1] == t[j-1]。此时，除了继承而来的dp[i-1][j]之外，还有dp[i-1][j-1]
	ns, nt := len(s), len(t)
	dp := make([][]int, ns+1)
	for i := range dp {
		dp[i] = make([]int, nt+1)
		dp[i][0] = 1
	}

	for i := 1; i <= ns; i++ {
		for j := 1; j <= nt; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[ns][nt]
}
