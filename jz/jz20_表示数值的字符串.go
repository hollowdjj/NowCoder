package jz

/*
请实现一个函数用来判断字符串str是否表示数值（包括科学计数法的数字，小数和整数）。

科学计数法的数字(按顺序）可以分成以下几个部分:
1.若干空格
2.一个整数或者小数
3.（可选）一个 'e' 或 'E' ，后面跟着一个整数(可正可负)
4.若干空格

小数（按顺序）可以分成以下几个部分：
1.若干空格
2.（可选）一个符号字符（'+' 或 '-'）
3. 可能是以下描述格式之一:
3.1 至少一位数字，后面跟着一个点 '.'
3.2 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
3.3 一个点 '.' ，后面跟着至少一位数字
4.若干空格

整数（按顺序）可以分成以下几个部分：
1.若干空格
2.（可选）一个符号字符（'+' 或 '-')
3. 至少一位数字
4.若干空格


例如，字符串["+100","5e2","-123","3.1416","-1E-16"]都表示数值。
但是["12e","1a3.14","1.2.3","+-5","12e+4.3"]都不是数值。
*/

func isNumeric(str string) bool {
	n := len(str)
	left, right := 0, n-1
	for left < right && str[left] == ' ' {
		left++
	}
	for left < right && str[right] == ' ' {
		right--
	}

	hasDot, hasNum, hasE, hasSign := false, false, false, false
	for left <= right {
		c := str[left]
		if c == '.' {
			//小数点符号只能出现一次，并且必须出现在符号E之前
			if hasDot || hasE {
				return false
			}
			hasDot = true
		} else if c == '+' || c == '-' {
			//符号前面不能有符号，数字以及小数点
			if hasSign || hasNum || hasDot {
				return false
			}
			hasSign = true
		} else if c == 'E' || c == 'e' {
			if hasE || !hasNum {
				return false
			}
			hasE = true
			hasDot, hasNum, hasSign = false, false, false
		} else if c <= '9' && c >= '0' {
			hasNum = true
		} else {
			return false
		}
		left++
	}
	return hasNum
}
