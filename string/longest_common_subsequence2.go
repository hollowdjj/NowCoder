package string

/*
最长公共子序列(二)

给定两个字符串str1和str2，输出两个字符串的最长公共子序列。如果最长公共子序列为空，则返回"-1"。
目前给出的数据，仅仅会存在一个最长的公共子序列
*/

func LCS2(s1, s2 string) string {
	//dp[i][j]表示s1[0:i]与s2[0:j]的最长公共子序列。因此，状态转移方程为：
	//1. s1[i-1] == s2[j-1]时，dp[i][j] = dp[i-1][j-1] + 1
	//2. s1[i-1] != s2[j-1]时，dp[i][j] = Max{dp[i-1][j],dp[i][j-1]}

	n1, n2 := len(s1), len(s2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	//选择
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	//逆过程得到解
	res := ""
	i, j := n1, n2
	for dp[i][j] > 0 {
		if s1[i-1] == s2[j-1] {
			res = string(s1[i-1]) + res
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	if res == "" {
		return "-1"
	}

	return res
}
