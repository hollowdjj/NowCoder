package string

/*
判断字符串是否是回文字符串

回文字符串是指字符串的正序和逆序完全一样
*/

func Palindromic(str string) bool {
	// write code here
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
