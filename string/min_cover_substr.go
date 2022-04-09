package string

import "math"

/*
最小覆盖子串
给出两个字符串 s 和 t，要求在 s 中找出最短的包含 t 中所有字符的连续子串。

例如：
S ="XDOYEZODEYXNZ" T ="XYZ"
找出的最短子串为"YXNZ".

注意：
如果 s 中没有包含 t 中所有字符的子串，返回空字符串 “”；
满足条件的子串可能有很多，但是题目保证满足条件的最短的子串唯一。
*/

func minWindow(s string, t string) string {
	n := len(t)
	need := make(map[byte]int)
	for i := 0; i < n; i++ {
		need[t[i]]++
	}

	n = len(s)
	count := len(need)
	left, right := 0, 0
	window := make(map[byte]int)
	minL := math.MaxInt32
	start := 0
	for right < n {
		for right < n && count > 0 {
			//增大滑动窗口直到窗口中包含了所有t中的字符
			c := s[right]
			window[c]++
			if window[c] == need[c] {
				count--
			}
			right++
		}
		//缩小滑动窗口到刚好能覆盖t中所有字符
		for left < right && count == 0 {
			c := s[left]
			if window[c] == need[c] {
				//不能再缩小了，更新结果
				if right-left < minL {
					minL = right - left
					start = left
				}
				count++
			}
			window[c]--
			left++
		}
	}

	if minL == math.MaxInt32 {
		return ""
	}
	return s[start : start+minL]
}
