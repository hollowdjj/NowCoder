package jz_II

/*
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从s 获得的 有效 IP 地址。
你可以按任何顺序返回答案。

有效IP地址正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
和 "192.168@1.1" 是无效IP地址。

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
*/

func restoreIpAddresses(s string) []string {
	n := len(s)
	res := make([]string, 0)
	var dfs func(curr []byte, i int, count int)
	dfs = func(curr []byte, i int, count int) {
		if i == n && count == 4 {
			res = append(res, string(curr[:len(curr)-1]))
			return
		}
		//剪枝
		if count == 4 || i >= n {
			return
		}
		num := 0

		//选一个
		num = num*10 + int(s[i]-'0')
		curr = append(curr, []byte{s[i], '.'}...)
		i++
		dfs(curr, i, count+1)
		if i >= n || num == 0 {
			return
		}
		//选两个
		num = num*10 + int(s[i]-'0')
		curr[len(curr)-1] = s[i]
		curr = append(curr, '.')
		i++
		dfs(curr, i, count+1)
		if i >= n {
			return
		}
		//选三个
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return
		}
		curr[len(curr)-1] = s[i]
		curr = append(curr, '.')
		i++
		dfs(curr, i, count+1)
	}
	dfs([]byte{}, 0, 0)
	return res
}
