package string

import (
	"strings"
	"unicode"
)

/*
字符串变形
对于一个长度为 n 字符串，我们需要对它做一些变形。

首先这个字符串中包含着一些空格，就像"Hello World"一样，然后我们要做的是把
这个字符串中由空格隔开的单词反序，同时反转每个字符的大小写。

比如"Hello World"变形后就变成了"wORLD hELLO"。

数据范围: 1≤n≤10^6, 字符串中包括大写英文字母、小写英文字母、空格。
进阶：空间复杂度 O(n)， 时间复杂度O(n)
*/

func Trans(s string, n int) string {
	builder := strings.Builder{}
	left, right := n-1, n-1
	for right >= 0 {
		if s[right] == ' ' {
			right--
			continue
		}
		//往前找到下一个空格
		left = right - 1
		for left >= 0 {
			if s[left] == ' ' {
				break
			}
			left--
		}
		builder.WriteString(LowAndUpTrans(s[left+1 : right+1]))
		if left > 0 {
			builder.WriteString(" ")
		}

		right = left
	}
	return builder.String()
}

func LowAndUpTrans(str string) string {
	n := len(str)
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		if unicode.IsLower(rune(str[i])) {
			res[i] = unicode.ToUpper(rune(str[i]))
		}
		if unicode.IsUpper(rune(str[i])) {
			res[i] = unicode.ToLower(rune(str[i]))
		}
	}
	return string(res)
}
