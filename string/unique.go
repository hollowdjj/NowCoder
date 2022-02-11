package string

/*
判断字符串中的字符是否都只出现过一次
*/

func IsUnique(str string) bool {
	dic := make(map[rune]bool)
	for _, r := range str {
		if _, ok := dic[r]; ok {
			return false
		}
		dic[r] = true
	}
	return true
}
