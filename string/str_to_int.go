package string

import (
	"math"
	"strings"
)

/*
把字符串转换成整数
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
传入的字符串可能有以下部分组成:
1.若干空格
2.（可选）一个符号字符（'+' 或 '-'）
3. 数字，字母，符号，空格组成的字符串表达式
4. 若干空格
*/

func StrToInt(s string) int {
	//去除前导空格和后导空格
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	//符号
	sign := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	s1 := []byte(s)
	res := 0
	for _, v := range s1 {
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')*sign
			if sign == 1 && res > math.MaxInt32 {
				return math.MaxInt32
			}
			if sign == -1 && res < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	return res
}
