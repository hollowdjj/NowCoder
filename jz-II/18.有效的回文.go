package jz_II

/*
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。

本题中，将空字符串定义为有效的 回文串 。

输入: s = "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
*/

func isPalindrome(s string) bool {
	n := len(s)
	left, right := 0, n-1
	for left < right {
		//找到左边第一个字母或数字
		for left < n && !isDigitOrAlpha(s[left]) {
			left++
		}
		//找到右边第一个字母或数字
		for right >= 0 && !isDigitOrAlpha(s[right]) {
			right--
		}
		if left < right {
			if !isEqual(s[left], s[right]) {
				return false
			}
		}
		left++
		right--
	}
	return true
}

func isDigitOrAlpha(s byte) bool {
	return s >= '0' && s <= '9' || s >= 'A' && s <= 'Z' || s >= 'a' && s <= 'z'
}

func isEqual(a, b byte) bool {
	if a >= '0' && a <= '9' {
		return a == b
	}
	return int(a-b) == 32 || int(b-a) == 32 || a == b
}
