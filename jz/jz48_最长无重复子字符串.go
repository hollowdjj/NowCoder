package jz

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

输入："abcabcbb"
输出：3
*/

func LongestNoRepeatedSubstring(s string) int {
	//滑动窗口
	n := len(s)
	res := 0
	window := make(map[byte]int)
	left, right := -1, 0
	for right < n {
		if i, ok := window[s[right]]; ok && i > left {
			left = i
		}
		res = max(res, right-left)
		window[s[right]] = right
		right++
	}
	return res
}
