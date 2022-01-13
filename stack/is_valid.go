package stack

import "nowcoder/utility"

/*
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

数据范围：字符串长度 0≤n≤10000
要求：空间复杂度 O(n)，时间复杂度 O(n)

输入："({})[]"
输出：true
*/

func IsValid(s string) bool {
	//使用一个辅助栈。一旦遇到左括号，那么入栈；一旦遇到右括号，若栈为空，返回false；若栈不为空。
	//弹出栈顶元素比较两者是否匹配即可
	var stack utility.Stack
	for _, c := range s {
		switch c {
		case '[', '{', '(':
			stack.Push(c)
		case ']':
			if stack.Empty() {
				return false
			}
			temp := (*stack.Pop()).(rune)
			if temp != '[' {
				return false
			}
		case '}':
			if stack.Empty() {
				return false
			}
			temp := (*stack.Pop()).(rune)
			if temp != '{' {
				return false
			}
		case ')':
			if stack.Empty() {
				return false
			}
			temp := (*stack.Pop()).(rune)
			if temp != '(' {
				return false
			}
		}
	}
	if !stack.Empty() {
		return false
	}
	return true
}
