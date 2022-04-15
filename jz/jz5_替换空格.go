package jz

/*
请实现一个函数，将一个字符串s中的每个空格替换成“%20”。
例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
*/

func replaceSpace(s string) string {
	n := len(s)
	res := ""
	left, right := 0, 0
	for right < n {
		//找到s[left:]中的第一个空格
		if s[right] == ' ' {
			res += s[left:right]
			res += "%20"
			left = right + 1
		}
		right++
	}

	if left < n {
		res += s[left:right]
	}
	return res
}
