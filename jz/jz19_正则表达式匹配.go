package jz

/*
请实现一个函数用来匹配包括'.'和'*'的正则表达式。
1.模式中的字符'.'表示任意一个字符
2.模式中的字符'*'表示它前面的字符可以出现任意次（包含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配
,但是与"aa.a"和"ab*a"均不匹配
*/

func match(str string, pattern string) bool {
	//dp[i][j]表示str[0:i]与pattern[0:j]是否匹配。对于'.'符号，由于其可以表示任意值
	//因此，不需要理会它，只需要关注'*'。
	//1. pattern[j-1] == '.'时，dp[i][j] = dp[i-1][j-1]
	//2. pattern[j-1] == 字母时，dp[i][j] = dp[i-1][j-1] && str[i-1] == str[j-1]
	//3. pattern[j-1] == '*'时，有两种情况：
	//3.1 '*'表示pattern[j-2]出现0次，那么dp[i][j] = dp[i][j-2]
	//3.2 '*'表示pattern[j-2]出现1次或多次时，那么dp[i][j] = dp[i-1][j] && (str[i-1] == p[j-2] || p[j-1] == '.')
	//除此之外，还需要注意base case：dp[0][j]
	ns, np := len(str), len(pattern)
	dp := make([][]bool, ns+1)
	for i := 0; i <= ns; i++ {
		dp[i] = make([]bool, np+1)
	}
	//base case
	dp[0][0] = true
	for j := 2; j <= np; j++ {
		if pattern[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= ns; i++ {
		for j := 1; j <= np; j++ {
			if pattern[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || dp[i-1][j] && (str[i-1] == pattern[j-2] || pattern[j-2] == '.')
			} else {
				dp[i][j] = dp[i-1][j-1] && str[i-1] == pattern[j-1]
			}
		}
	}
	return dp[ns][np]
}
