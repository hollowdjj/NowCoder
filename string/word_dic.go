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
