package string

/*
括号生成

给出n对括号，请编写一个函数来生成所有的由n对括号组成的合法组合。
例如，给出n=3，解集为：
"((()))", "(()())", "(())()", "()()()", "()(())"
*/

func GenerateParenthesis(n int) []string {
	//当左括号的数量小于n时，一直添加左括号。
	//当右括号的数量小于左括号的数量时，添加右括号
	res := make([]string, 0)
	var dfs func(curr string, left, right int)
	dfs = func(curr string, left, right int) {
		if left == n && right == n {
			res = append(res, curr)
			return
		}
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
