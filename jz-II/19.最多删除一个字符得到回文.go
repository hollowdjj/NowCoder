package jz_II

/*
给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。

输入: s = "abca"
输出: true
解释: 可以删除 "c" 字符 或者 "b" 字符
*/

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			//判断s[left+1,right]是否为回文，即删除s[left]
			valid1 := valid(s[left+1 : right+1])
			//判断s[left,right-1]是否为回文，即删除s[right]
			valid2 := valid(s[left:right])
			return valid1 || valid2
		}
	}
	return true
}

func valid(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
