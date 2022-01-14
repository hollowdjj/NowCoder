package stack

import (
	"nowcoder/utility"
	"unicode"
)

/*
给定一个字符串形式的表达式 s ，请你实现一个计算器并返回结果。

数据范围：表达式长度满足  ，字符串中包含 + , - , * , / , ( , ) ，保证表达式合法。
*/

func Calculate2(s string) int {
	var nums utility.Stack
	number := 0
	sign := '+'
	n := len(s)
	for i := 0; i < n; i++ {
		c := rune(s[i])
		if c == ' ' {
			continue
		}
		if unicode.IsDigit(c) {
			//如果c是数字的话，那么递归得到完整的数字(例如356)
			number = number*10 + int(c-'0')
		}
		if c == '(' {
			//如果c是左括号的话， 那么一直往后找到匹配的右括号
			count := 1
			j := i + 1
			for count > 0 {
				if s[j] == '(' {
					count++
				} else if s[j] == ')' {
					count--
				}
				j++
			}
			//循环结束后，j是与c匹配的右括号的下一个位置
			number = Calculate2(s[i+1 : j-1]) //递归计算括号中的子表达式
			i = j - 1                         //i需要等于与c匹配的右括号的索引
		}
		if !unicode.IsDigit(c) || i == n-1 {
			switch sign {
			case '*':
				nums.Push((*nums.Pop()).(int) * number)
			case '/':
				nums.Push((*nums.Pop()).(int) / number)
			case '+':
				nums.Push(number)
			case '-':
				nums.Push(-1 * number)
			}
			sign = c
			number = 0
		}
	}
	sum := 0
	for !nums.Empty() {
		sum += (*nums.Pop()).(int)
	}
	return sum
}
