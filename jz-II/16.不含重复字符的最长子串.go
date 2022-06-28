package jz_II

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。请注意，你的答案必须是子串的长度
"pwke"是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	dic := make(map[byte]int) //滑动窗口中的字符及其出现的次数
	left, right := 0, 0
	n, res := len(s), 0
	for right < n {
		dic[s[right]]++
		for dic[s[right]] > 1 {
			dic[s[left]]--
			left++
		}
		if t := right - left + 1; t > res {
			res = t
		}
		right++
	}
	return res
}

func lengthOfLongestSubstring1(s string) int {
	dic := make(map[byte]int) //字符出现的最新位置
	left, right := -1, 0
	n, res := len(s), 0
	for right < n {
		if index, ok := dic[s[right]]; ok {
			//缩小滑动窗口
			if index > left {
				left = index
			}
		}
		dic[s[right]] = right
		if right-left > res {
			res = right - left
		}
		right++
	}
	return res
}
