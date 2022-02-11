package string

import "strings"

/*
大数加法
*/

func BigNumPlus(s string, t string) string {
	i, j := len(s)-1, len(t)-1

	b := strings.Builder{}
	carry := 0
	for i >= 0 && j >= 0 {
		val := int(s[i]-'0'+t[j]-'0') + carry
		carry = val / 10
		val %= 10
		b.WriteRune(int32(val) + '0')
		i--
		j--
	}

	for i >= 0 {
		val := int(s[i]-'0') + carry
		carry = val / 10
		val %= 10
		b.WriteRune(int32(val) + '0')
		i--
	}

	for j >= 0 {
		val := int(t[j]-'0') + carry
		carry = val / 10
		val %= 10
		b.WriteRune(int32(val) + '0')
		j--
	}

	if carry > 0 {
		b.WriteRune(int32(carry) + '0')
	}

	return Reverse(b.String())
}
