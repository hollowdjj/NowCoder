package string

/*
最长无重复子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
数据范围:s.length≤40000

例如："abcabcbb"
返回值：3
说明：因为无重复字符的最长子串是"abc"，所以其长度为 3。
*/

func LongestNoRepeatSubstr(s string) int {
	//双指针加哈希表。哈希表保存的是每个字符最新出现的位置
	n := len(s)
	dic := make(map[byte]int)
	left, right, res := -1, 0, 1
	for right < n {
		if index, ok := dic[s[right]]; ok {
			//只有当重复的元素上一次出现的位置在left之后时，才更新left
			//否则会发生向前更新的错误。例如：pwwabcwfw。
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
