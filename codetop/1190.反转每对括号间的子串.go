package codetop

/*
给出一个字符串 s（仅含有小写英文字母和括号）。
请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
注意，您的结果中 不应 包含任何括号。

输入：s = "(u(love)i)"
输出："iloveu"
解释：先反转子字符串 "love" ，然后反转整个字符串。
*/

func reverseParentheses(s string) string {
	var str []byte
	var stack [][]byte //当前字符串
	for i := range s {
		if s[i] == '(' {
			//遇到右括号就将str压栈，并清空str
			stack = append(stack, str)
			str = []byte{}
		} else if s[i] == ')' {
			//遇到右括号，那么需要将里面的字符串，也就是str翻转
			left, right := 0, len(str)-1
			for left < right {
				str[left], str[right] = str[right], str[left]
				left++
				right--
			}
			//然后合并栈顶字符串，栈顶字符串其实就是右括号前面的字符串
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			str = append(top, str...)
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
