package string

/*
判断子序列
给定两个字符串 S 和 T ，判断 S 是否是 T 的子序列。
即是否可以从 T 删除一些字符转换成 S。

数据范围：1≤len(S)≤100，1≤len(T)≤10^4，保证字符串中仅含有小写字母
例如："nowcoder","nowcoder" 输出： true

*/

func IsSubsequence(S string, T string) bool {
	//双指针法。i,j分别指向S和T的头。如果S[i] == T[j]，那么i和j都加1
	//否则只有j加1。最后若i等于len(S)说明存在这样的子序列
	ns, nt := len(S), len(T)
	i, j := 0, 0
	for i < ns && j < nt {
		if S[i] == T[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == ns {
		return true
	}
	return false
}
