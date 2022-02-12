package string

/*
找到字符串中的最长回文子串长度

例如："ababc" 返回3
说明：最长的回文子串为"aba"与"bab"，长度都为3
*/
func LongestPalindrome(A string) int {
	//dp[i][j]表示A[i...j]子串是否是一个回文字符串。因此，我们可以得到状态
	//转移方程： dp[i][j] = dp[i+1][j-1] A[i] = A[j]且 j-i>=2
	//			dp[i][j] = true A[i] = A[j]且0<=j-i< 2
	n := len(A)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true //单个字符也是一个回文字符子串
	}

	res := 1 //回文子串的长度至少为1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j-i >= 2 {
				if A[i] == A[j] && dp[i+1][j-1] {
					dp[i][j] = true
					if j-i+1 > res {
						res = j - i + 1
					}
				}
			} else {
				if A[i] == A[j] {
					dp[i][j] = true
					if res < 2 {
						res = 2
					}
				}
			}
		}
	}
	return res
}
