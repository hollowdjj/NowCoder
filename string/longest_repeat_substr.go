package string

/*
最长重复子串
定义重复字符串是由两个相同的字符串首尾拼接而成，例如 abcabcabcabc 便是长度为6
的一个重复字符串，而 abcbaabcba 则不存在重复字符串。
给定一个字符串，请返回其最长重复子串的长度。若不存在任何重复字符子串，则返回 0 。
*/
func LongestRepeatSubstr(a string) int {
	n := len(a)
	max := 0
	for i := 0; i < n-1; i++ { //起点
		for j := 1; j < n-i; j++ { //子串长度
			if j%2 == 1 {
				continue
			}
			if check(i, i+j-1, a) {
				if j > max {
					max = j
				}
			}
		}
	}
	return max
}

func check(i, j int, s string) bool {
	mid := i + (j-i)/2
	for k := i; k <= mid; k++ {
		if s[k] != s[mid+k-i+1] {
			return false
		}
	}
	return true
}
