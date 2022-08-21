package codetop

import "strings"

/*
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次
数都不少于 k 。返回这一子串的长度。

输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
*/

func longestSubstring(s string, k int) int {
	if s == "" {
		return 0
	}
	//如果有一个字符出现的次数少于k，那么所有包含了该字符的字符串都不满足要求
	//因此，可以以该字符分割字符串，然后递归求解。
	dic := [26]int{}
	for i := range s {
		dic[int(s[i]-'a')]++
	}
	var b byte
	for i, count := range dic {
		if count > 0 && count < k {
			b = 'a' + byte(i)
			break
		}
	}
	if b == 0 {
		return len(s)
	}
	res := 0
	strs := strings.Split(s, string(b))
	for _, sub := range strs {
		res = max(res, longestSubstring(sub, k))
	}
	return res
}
