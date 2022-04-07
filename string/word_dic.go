package string

/*
单词拆分(一)

给定一个字符串和一个字符串数组，在字符串的任意位置拆分任意次后得到的字符串集合是否是给定字符串数组的子集。

输入："nowcodernow",["now","coder"]
输出：true
说明：可拆分得到“now” “code” “now”。出现了两次"now"，但认为是一个

*/
func WordDiv(s string, strs []string) bool {
	dic := make(map[string]bool)
	for _, str := range strs {
		dic[str] = true
	}

	n := len(s)
	res := false
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = true
			return
		}
		if !res {
			for j := i + 1; j <= n && !res; j++ {
				if dic[s[i:j]] {
					dfs(j)
				}
			}
		}
	}
	dfs(0)
	return res
}

func WordDivDp(s string, strs []string) bool {
	//用回溯做，会超时。考虑使用动态规划。dp[i]表示s[0:i]能否拆分出字典中的单词
	//dp[i] = true when dp[j] == true && strs.Contains(s[j:i]), where 0<=j<=i-1
	dic := make(map[string]bool)
	for _, str := range strs {
		dic[str] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			prefix := s[j:i]
			//dp[j]其实就代表了s[0:j]的结果。因此，这里的意思其实就是s[0:j]和s[j:i]
			//如果都是字典中的字符串的话，dp[i]=true。其实就是遍历可能的拆分点，拆分
			//s[0:i]。
			if dic[prefix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
