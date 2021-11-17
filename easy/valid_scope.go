package easy

import "fmt"

/*
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

要求：空间复杂度 O(n)，时间复杂度 O(n)
*/

//空间复杂度为O(1)不满足题目要求
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	left, right := n/2-1, n/2
	for i := 0; i < n/2; i++ {
		if (s[left] == '[' && s[right] == ']') ||
			(s[left] == '{' && s[right] == '}') ||
			(s[left] == '(' && s[right] == ')') {
			continue
		}

		return false
	}

	return true
}

func IsValidByStack(s string) bool {
	/*
		要想达到题目中空间复杂度为O(n)的要求，需要使用一个辅助栈。具体的做法为：
		当遇到左括号时，将其压入栈中。一旦遇到右括号，那么弹出栈顶元素，判断栈顶
		是否匹配。
	*/
	stack := StackRune([]rune{})
	for _, c := range s {
		switch c {
		case '[', '{', '(':
			stack.Push(c)
		case ']':
			if stack.Empty() || stack.Top() != '[' {
				return false
			}
			stack.Pop()
		case '}':
			if stack.Empty() || stack.Top() != '{' {
				return false
			}
			stack.Pop()
		case ')':
			if stack.Empty() || stack.Top() != '(' {
				return false
			}
			stack.Pop()
		}
	}
	return true
}

func TestIsValid() {
	s := "[{()}]"
	fmt.Println(IsValidByStack(s))
}

type StackRune []rune

func (s *StackRune) Push(r rune) {
	*s = append(*s, r)
}

func (s *StackRune) Pop() rune {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *StackRune) Top() rune {
	return (*s)[len(*s)-1]
}

func (s *StackRune) Empty() bool {
	if len(*s) > 0 {
		return false
	}
	return true
}
