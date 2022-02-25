package string

/*
最长公共子串
给定两个字符串str1和str2,输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一。

输入："1AB2345CD","12345EF"
输出："2345"
*/

func LongestCommonSubstr(str1, str2 string) string {
	//dp[i][j]表示str1中以str1[i-1]结尾以及str2中以str2[j-1]结尾的最长公共子串的长度。
	//故，状态转移方程：
	//1. dp[i][j] = 0 					when str1[i-1] != str2[j-1]
	//2. dp[i][j] = dp[i-1][j-1] + 1    when str1[i-1] == str2[j-1]
	n1, n2 := len(str1), len(str2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	maxLength := 0
	end := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					end = i
				}
			}
		}
	}

	return str1[end-maxLength : end]
}
