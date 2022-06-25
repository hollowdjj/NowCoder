package jz_II

/*
正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且有效的括号组合。

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/
func generateScopes(n int) []string {
	res := make([]string, 0)
	var dfs func(curr string, left, right int)
	dfs = func(curr string, left, right int) {
		if left == n && right == n {
			res = append(res, curr)
			return
		}
		//左括号数量小于n时，添左括号；右括号数量小于左括号时，添加右括号。
		if left < n {
			dfs(curr+"(", left+1, right)
		}
		if right < left {
			dfs(curr+")", left, right+1)
		}
	}
	dfs("", 0, 0)
	return res
}
