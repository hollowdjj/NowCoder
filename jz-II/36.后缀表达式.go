package jz_II

import "strconv"

/*
根据 逆波兰表示法，求该后缀表达式的计算结果。
有效的算符包括+、-、*、/。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：
整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
*/

func evalRPN(tokens []string) int {
	//后缀表达式一定是从前往后看的！！！！！
	stack := make([]int, 0)
	for _, s := range tokens {
		if s[len(s)-1] >= '0' && s[len(s)-1] <= '9' {
			v, _ := strconv.Atoi(s)
			stack = append(stack, v)
		} else {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			res := 0
			stack = stack[:len(stack)-2]
			if s == "+" {
				res = b + a
			} else if s == "-" {
				res = b - a
			} else if s == "*" {
				res = b * a
			} else {
				res = b / a
			}
			stack = append(stack, res)
		}
	}
	return stack[len(stack)-1]
}
