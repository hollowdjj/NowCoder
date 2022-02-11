package stack

import (
	"strings"
	"unicode"
)

/*
给一个加密过的字符串解码，返回解码后的字符串。

加密方法是：k[c] ，表示中括号中的 c 字符串重复 k 次，例如 3[a] 解码结果是 aaa ，保证输入字符串符合规则。不会出现类似 3a , 3[3] 这样的输入。
*/

func DecodeString(s string) string {
	var str string
	number := 0
	builder := strings.Builder{}
	n := len(s)
	for i := 0; i < n; i++ {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			//一旦遇到数字，那么循环得到完整的数字
			number = number*10 + int(c-'0')
		}
		if c == '[' {
			//一直往后找到匹配的右括号
			count := 1
			j := i + 1
			for count > 0 {
				if s[j] == '[' {
					count++
				} else if s[j] == ']' {
					count--
				}
				j++
			}
			str = DecodeString(s[i+1 : j-1])
			i = j - 1
		}
		if !unicode.IsDigit(c) || i == n-1 {
			if number != 0 {
				str = strings.Repeat(str, number)
				builder.WriteString(str)
			} else {
				builder.WriteRune(c)
			}
			number = 0
			str = ""
		}
	}
	return builder.String()
}
