package string

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

func MinWindow(S string, T string) string {
	ns, nt := len(S), len(T)
	if ns < nt || nt == 0 {
		return ""
	}
	//先使用一张哈希表记录T中的所有字符出现的次数
	dic := make(map[byte]int)
	count := 0
	for i := 0; i < nt; i++ {
		if dic[T[i]] == 0 {
			count++
		}
		dic[T[i]]++
	}

	//使用滑动窗口遍历字符串S
	left, right := 0, 0
	start, end := -1, ns
	window := make(map[byte]int)
	match := 0
	for right < ns {
		c := S[right]
		if dic[c] > 0 {
			window[c]++
			if window[c] == dic[c] {
				//说明对于T中的所有字符c都能够在S[left,right]中找到，匹配度加1
				match++
			}
		}
		right++
		//这里一定是for循环。因为left++不一定会导致match--。所以我们要一直更新解，直到math != count
		for match == count {
			//T中的所有字符已经在S[left,right]中找到，更新解
			if right-left < end-start {
				start, end = left, right
			}
			c1 := S[left]
			if dic[c1] > 0 {
				window[c1]--
				if window[c1] < dic[c1] {
					match--
				}
			}
			left++
		}
	}
	if start == -1 {
		return ""
	}
	return S[start:end]
}
