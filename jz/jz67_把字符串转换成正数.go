package jz

import "math"

/*
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。传入的字符串可能有以下部分组成:
1.若干空格
2.（可选）一个符号字符（'+' 或 '-'）
3. 数字，字母，符号，空格组成的字符串表达式
4. 若干空格

转换算法如下:
1.去掉无用的前导空格
2.第一个非空字符为+或者-号时，作为该整数的正负号，如果没有符号，默认为正数
3.判断整数的有效部分：
3.1 确定符号位之后，与之后面尽可能多的连续数字组合起来成为有效整数数字，如果没有有效的整数部分，那么直接返回0
3.2 将字符串前面的整数部分取出，后面可能会存在存在多余的字符(字母，符号，空格等)，这些字符可以被忽略，它们对于函数不应该造成影响
3.3  整数超过 32 位有符号整数范围 [−2^31,  2^31 − 1]，需要截断这个整数，使其保持在这个范围内
4.去掉无用的后导空格
*/

func StrToInt(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	//去除前导空格
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	//得到符号位
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	} else if s[i] < '0' || s[i] > '9' {
		return 0
	}

	//得到数字
	res := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		//fmt.Println(res)
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}

	return res * sign
}
