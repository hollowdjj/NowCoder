package jz_II

/*
给定三个字符串s1、s2、s3，请判断s3能不能由s1和s2交织（交错）组成。

两个字符串 s 和 t 交织的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	//dp[i][j]表示s1[0:i]和s2[0:j]能否交织组成s3[0:i+j]
	//状态转移方程为：
	//1. 若s1[i-1] == s3[i+j-1]，那么dp[i][j] = dp[i-1][j]
	//2. 若s2[j-1] == s3[i+j-1]，那么dp[i][j] = dp[i][j-1]
	n1, n2 := len(s1), len(s2)
	if n1+n2 != len(s3) {
		return false
	}
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || s1[i-1] == s3[i+j-1] && dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || s2[j-1] == s3[i+j-1] && dp[i][j-1]
			}
		}
	}
	return dp[n1][n2]
}
