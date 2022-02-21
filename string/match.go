package string

/*
正则表达式匹配
请实现一个函数用来匹配包括'.'和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（包含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab*a"均不匹配

数据范围:
1.str 可能为空，且只包含从 a-z 的小写字母。
2.pattern 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。
3. 0 <= str.length <= 20
4. 0 <= pattern.length <= 30
*/

func Match(str, pattern string) bool {
	//dp[i][j]表示str[0:i]与pattern[0:j]是否匹配。在pattern字符串中只有字母，'.'以及'*'，故状态转移方程
	//也是针对这三种情况进行讨论：
	//1. 当p[j-1] = 字母时，显然dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
	//2. 当p[j-1] = '.'时，很显然dp[i][j] = dp[i-1][j-1]。因为'.'表示任意字符
	//3. 当p[j-1] = '*'时：
	//(3.1)*表示0个字符时，dp[i][j] = dp[i][j-2]
	//(3.2)*匹配多个字符时，dp[i][j] = dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] = '.')

	//初始化dp数组以及base case
	ns, np := len(str), len(pattern)
	dp := make([][]bool, ns+1)
	for i := 0; i <= ns; i++ {
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true
	for j := 2; j <= np; j++ {
		//诸如a*a*a*a*a*
		if pattern[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	//选择
	for i := 1; i <= ns; i++ {
		for j := 1; j <= np; j++ {
			if pattern[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern[j-1] == '*' {
				dp[i][j] = j >= 2 && dp[i][j-2] || dp[i-1][j] && (str[i-1] == pattern[j-2] || pattern[j-2] == '.')
			} else {
				dp[i][j] = dp[i-1][j-1] && str[i-1] == pattern[j-1]
			}
		}
	}

	return dp[ns][np]
}
