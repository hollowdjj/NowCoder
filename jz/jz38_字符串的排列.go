package jz

/*
输入一个长度为 n 字符串，打印出该字符串中字符的所有排列，你可以以任意顺序返回这个字符串数组。
例如输入字符串ABC,则输出由字符A,B,C所能排列出来的所有字符串ABC,ACB,BAC,BCA,CBA和CAB。

输入："aab"
输出：["aab","aba","baa"]
*/

func PermutationString(str string) []string {
	//有重复元素的全排列问题
	n := len(str)
	used := make([]bool, n)
	res := make([]string, 0)
	var backtrack func(curr []byte)
	backtrack = func(curr []byte) {
		if len(curr) == n {
			res = append(res, string(curr))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] || i > 0 && str[i] == str[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			curr = append(curr, str[i])
			backtrack(curr)
			curr = curr[:len(curr)-1]
			used[i] = false
		}
	}
	backtrack([]byte{})
	return res
}
