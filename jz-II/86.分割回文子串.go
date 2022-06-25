package jz_II

/*
给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

输入：s = "google"
输出：[["g","o","o","g","l","e"],["g","oo","g","l","e"],["goog","l","e"]]
*/

func partition(s string) [][]string {
	n := len(s)
	res := make([][]string, 0)
	dic := make(map[string]bool)
	var dfs func(curr []string, i int)
	dfs = func(curr []string, i int) {
		if i == n {
			temp := make([]string, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for j := i; j < n; j++ {
			str := s[i : j+1]
			_, ok := dic[str]
			if !ok {
				dic[str] = isPal(str)
			}
			if dic[str] {
				curr = append(curr, str)
				dfs(curr, j+1)
				curr = curr[:len(curr)-1]
			}
		}
	}
	dfs([]string{}, 0)
	return res
}

func isPal(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
