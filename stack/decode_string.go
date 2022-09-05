package stack

import (
	"strings"
)

/*
给一个加密过的字符串解码，返回解码后的字符串。

加密方法是：k[c] ，表示中括号中的 c 字符串重复 k 次，例如 3[a] 解码结果是 aaa ，保证输入字符串符合规则。不会出现类似 3a , 3[3] 这样的输入。
*/

func DecodeString(s string) string {
	i, n := 0, len(s)
	res := ""
	for ; i < n; i++ {
		if s[i] <= '9' && s[i] >= '0' {
			num := 0
			//找到完整数字
			for ; i < n && s[i] != '['; i++ {
				num = num*10 + int(s[i]-'0')
			}
			//递归求解括号内的字符串
			count := 0
			left := i + 1
			for ; i < n; i++ {
				if s[i] == '[' {
					count++
				} else if s[i] == ']' {
					count--
				}
				if count == 0 {
					break
				}
			}
			//fmt.Println(num)
			str := DecodeString(s[left:i])
			//fmt.Println(str)
			res += strings.Repeat(str, num)
		} else {
			res += string(s[i])
		}
	}
	return res
}
