package string

import "strings"

/*
单词拆分(二)
给定一个字符串 s 和一个字符串数组 dic ，在字符串 s 的任意位置添加任意多个空格后得到的
字符串集合是给定字符串数组dic的子集(即拆分后的字符串集合中的所有字符串都在 dic 数组中)，
你可以以任意顺序 返回所有这些可能的拆分方案。

数据范围：字符串长度满足1≤n≤20，数组长度满足1≤n≤10，数组中字符串长度满足1≤n≤20

输入："nowcoder",["now","coder","no","wcoder"]
输出：["no wcoder","now coder"]
*/

func WordDiv2(s string, dic []string) []string {
	//题目的意思是，想办法把字符串s进行拆分，使得拆分出来的每个子串都能够在
	//字符串集合dic中找到。这道题可以考虑利用dfs来解决。即固定1位往后递归
	m := make(map[string]bool)
	for _, str := range dic {
		m[str] = true
	}

	curr := make([]string, 0)
	res := make([]string, 0)
	var dfs func(i int)
	dfs = func(i int) {
		n := len(s)
		if i == n {
			res = append(res, strings.Join(curr, " "))
			return
		}
		for j := i + 1; j <= n; j++ {
			if m[s[i:j]] {
				//只有当子串s[i:j]出现在了dic中时，才往后进行搜索。并且是dfs(j)
				//意味着下一次递归是要固定s[j]。
				curr = append(curr, s[i:j])
				dfs(j)
				curr = curr[:len(curr)-1]
			}
		}
	}
	dfs(0)
	return res
}
