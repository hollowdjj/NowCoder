package string

/*
最长公共子序列(一)

给定两个字符串 s1 和 s2，长度为m和n 。求两个字符串最长公共子序列的长度。
所谓子序列，指一个字符串删掉部分字符（也可以不删）形成的字符串。例如：字符串 "arcaea" 的子序列有 "ara" 、 "rcaa" 等。
但 "car" 、 "aaae" 则不是它的子序列。
所谓 s1 和 s2 的最长公共子序列，即一个最长的字符串，它既是 s1 的子序列，也是 s2 的子序列。
数据范围 : 0≤m,n≤1000 。保证字符串中的字符只有小写字母。

输入："abcde","bdcaaa"
输出：2
说明：最长公共子序列为 "bc" 或 "bd" ，长度为2
*/

func LCS(s1, s2 string) int {
	//dp[i][j]表示s1[0:i]和s2[0:j]最长公共子序列的长度，那么状态转移方程为：
	//1.当s1[i-1]==s2[j-1]时，dp[i][j] = dp[i-1][j-1]+1
	//2.当s1[i-1]!=s2[j-1]时，dp[i][j] = Max{dp[i-1][j],dp[i][j-1]}
	n1, n2 := len(s1), len(s2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}
