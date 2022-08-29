package jz

import "strings"

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

func isNumber(s string) bool {
	//去掉首尾空格
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	hasNum, hasDot, hasSign, hasE := false, false, false, false
	for _, ch := range s {
		if ch == ' ' {
			//空格只能出现在首尾
			return false
		} else if ch == '.' {
			//'.'只能出现一次，并且只能出现在E后面
			if hasDot || hasE {
				return false
			}
			hasDot = true
		} else if ch == 'e' || ch == 'E' {
			//'E'只能出现一次，并且E前面必须有数字
			if hasE || !hasNum {
				return false
			}
			//'E'出现后，要把hasNum,hasSign,hasDot置为false
			//这是因为E后面必须有数字（hasNum置为false），可能有符号(hasSign置为false)
			//hasDot置为false是因为判断有无字符时，hasNum || hasSign || hasDot
			hasE = true
			hasNum, hasSign, hasDot = false, false, false
		} else if ch <= '9' && ch >= '0' {
			hasNum = true
		} else if ch == '+' || ch == '-' {
			//+ - 前面不能有数组、符号以及小数点
			if hasNum || hasSign || hasDot {
				return false
			}
			hasSign = true
		} else {
			return false
		}
	}
	//必须有数字
	return hasNum
}
