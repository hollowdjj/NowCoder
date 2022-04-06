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

func MinWindow(s string, t string) string {
	need := make(map[byte]int)
	nt, ns := len(t), len(s)
	for i := 0; i < nt; i++ {
		need[t[i]]++
	}
	count := len(need)
	left, right := 0, 0 //滑动窗口左右指针
	start, length := 0, math.MaxInt32
	window := make(map[byte]int)
	for right < ns {
		//滑动窗口right指针一直右移，直到滑动窗口中包含了所有t中的字符
		for right < ns {
			c := s[right]
			window[c]++
			if window[c] == need[c] {
				count--
				if count == 0 {
					break
				}
			}
			right++
		}

		if count == 0 {
			//往右移动滑动窗口的左指针，直到滑动窗口中的元素刚好覆盖t中的字符，然后更新结果
			stop := false
			for left <= right && !stop {
				c := s[left]
				if window[c] == need[c] {
					//窗口无法缩小了，更新答案
					if right-left+1 < length {
						start = left
						length = right - left + 1
					}
					stop = true
					count++
				}
				window[c]--
				left++
			}
		}
		right++
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : length+start]
}
