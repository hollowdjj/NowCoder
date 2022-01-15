package stack

/*
给定一个字符串s，字符串s只包含以下三种字符: (，*，)，请你判断 s是不是一个合法的括号字符串。合法括号字符串有如下规则:
1.左括号'('必须有对应的右括号')'
2.右括号')'必须有对应的左括号'('
3.左括号必须在对应的右括号前面
4.*可以视为单个左括号，也可以视为单个右括号，或者视为一个空字符
5.空字符串也视为合法的括号字符串
*/

func IsValidString(s string) bool {
	//将*当成左括号正序遍历一遍，然后再将*当成右括号逆序遍历一遍
	//从左往右遍历时，左括号的数量必须时刻大于等于右括号的数量，否则不合法
	//从右往左遍历时，右括号的数量必须时刻大于等于左括号的数量，否则不合法
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '*' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if left < right {
			return false
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || s[i] == '*' {
			right++
		} else if s[i] == '(' {
			left++
		}
		if right < left {
			return false
		}
	}
	return true
}
