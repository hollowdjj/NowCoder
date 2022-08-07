package jz_II

/*
给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存在符合
条件的子字符串，则返回空字符串 ""。

如果 s 中存在多个符合条件的子字符串，返回任意一个。

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最短子字符串 "BANC" 包含了字符串 t 的所有字符 'A'、'B'、'C'
*/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	dic := make(map[byte]int)
	for i := range t {
		dic[t[i]]++
	}
	win := make(map[byte]int)
	res := ""
	left, right := 0, 0
	count := 0
	for right < len(s) {
		b := s[right]
		win[b]++
		if win[b] == dic[b] {
			count++
		}
		//缩小滑动窗口，找最短子串
		for count == len(dic) {
			if res == "" || len(res) > right-left+1 {
				res = s[left : right+1]
			}
			win[s[left]]--
			if win[s[left]] < dic[s[left]] {
				count--
			}
			left++
		}
		right++
	}

	return res
}
