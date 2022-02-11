package string

import (
	"fmt"
	"strings"
)

/*
压缩字符串
利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2bc5a3。
1.如果只有一个字符，1不用写
2.字符串中只包含大小写英文字母（a至z）。
数据范围:
0<=字符串长度<=50000
要求：时间复杂度O(N）
*/

func CompressString(param string) string {
	n := len(param)
	if n <= 1 {
		return param
	}
	builder := strings.Builder{}
	prev := param[0]
	count := 1
	for i := 1; i < n; i++ {
		if param[i] == prev {
			count++
		} else {
			builder.WriteByte(prev)
			if count > 1 {
				builder.WriteString(fmt.Sprintf("%d", count))
			}
			prev = param[i]
			count = 1
		}
	}
	//这里要处理尾巴上的字符
	builder.WriteByte(param[n-1])
	if count > 1 {
		builder.WriteString(fmt.Sprintf("%d", count))
	}
	return builder.String()
}
