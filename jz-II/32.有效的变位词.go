package jz_II

/*
给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
注意：若s和t中每个字符出现的次数都相同且字符顺序不完全相同，则称s和t互为变位词（字母异位词）。

输入: s = "anagram", t = "nagaram"
输出: true
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	dic := make(map[rune]int)
	for _, r := range s {
		dic[r]++
	}
	for _, r := range t {
		dic[r]--
		if dic[r] < 0 {
			return false
		}
	}
	return true
}
