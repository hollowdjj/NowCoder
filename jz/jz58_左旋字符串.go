package jz

/*
汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。
对于一个给定的字符序列  S ，请你把其循环左移 K 位后的序列输出。例如，字符序列 S = ”abcXYZdef”
要求输出循环左移 3 位后的结果，即 “XYZdefabc”
*/

func LeftRotateString(str string, n int) string {
	if len(str) == 0 {
		return str
	}
	index := n % len(str)
	return str[index:] + str[:index]
}
