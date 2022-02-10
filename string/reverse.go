package string

/*
反转字符串
*/

func Reverse(str string) string {
	b := []byte(str)
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return string(b)
}
