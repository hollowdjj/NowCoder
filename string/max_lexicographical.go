package string

import "strings"

/*
二进制取反
有一个二进制字符串num，可以选择该串中的任意一段区间进行取反(可以进行一次或不进行)，
取反指将0变为1，将1变为0。那么取反之后的num可能的最大的字典序是多少呢。
例如：011100 取反后 是111100(注意，只能对一段区间进行取反)
*/

func MaxLexicographical(num string) string {
	//其实就是找到高位的0序列，将这个序列全变为1即可
	builder := strings.Builder{}
	n := len(num)
	for i := 0; i < n; i++ {
		if num[i] == '0' {
			//一直往下找到第一个不是0的位置
			j := i + 1
			for ; j < n; j++ {
				if num[j] != '0' {
					break
				}
			}
			for k := i; k < j; k++ {
				builder.WriteRune('1')
			}
			builder.WriteString(num[j:])
			break
		} else {
			builder.WriteByte(num[i])
		}
	}
	return builder.String()
}
