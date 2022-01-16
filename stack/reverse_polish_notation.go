package stack

import (
	"nowcoder/utility"
	"strconv"
)

/*
逆波兰表达式求值:
在平时，我们习惯将表达式写成(1+2)*(3+4)，加减乘等运算符写在中间，因此称呼为中缀表达式
而波兰表达式的写法为：(*(+1 2)(+ 3 4))   去掉括号为* + 1 2 + 3 4，运算符在前面，也称为前缀表达式
逆波兰表达式的写法为: ((1 2 +)(3 4 +)*)，去掉括号为1 2 + 3 4 + *，运算符在后面，也称为后缀表达式
*/

func EvalReversePolishNotation(tokens []string) int {
	//逆波兰表达式的求解是比较简单的。当遇到数字时，直接压栈。遇到运算符时，取栈顶两元素
	//计算值然后压栈。需要注意的是，由于是栈所以计算的时候b在前而a在后
	var s utility.Stack
	for _, str := range tokens {
		switch str {
		case "+":
			a := (*s.Pop()).(int)
			b := (*s.Pop()).(int)
			s.Push(b + a)
		case "-":
			a := (*s.Pop()).(int)
			b := (*s.Pop()).(int)
			s.Push(b - a)
		case "*":
			a := (*s.Pop()).(int)
			b := (*s.Pop()).(int)
			s.Push(b * a)
		case "/":
			a := (*s.Pop()).(int)
			b := (*s.Pop()).(int)
			s.Push(b / a)
		default:
			num, _ := strconv.Atoi(str)
			s.Push(num)
		}
	}
	return (*s.Pop()).(int)
}
