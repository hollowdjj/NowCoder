package string

/*
数字字符串转换成IP地址
现在有一个只包含数字的字符串，将该字符串转化成IP地址的形式，返回所有可能的情况。
例如：
给出的字符串为"25525522135",
返回["255.255.22.135", "255.255.221.35"]. (顺序没有关系)

数据范围：字符串长度 0≤n≤12
要求：空间复杂度O(n!),时间复杂度O(n!)

注意：ip地址是由四段数字组成的数字序列，格式如 "x.x.x.x"，其中 x 的范围应当是 [0,255]。
*/

var ipRes []string

func RestoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return nil
	}
	helper([]byte{}, 0, 0, s)
	return ipRes
}

func helper(curr []byte, i, j int, s string) {
	//递归终止条件
	n := len(s)
	if i == 4 && j == n {
		ipRes = append(ipRes, string(curr[:len(curr)-1]))
		return
	}
	if i == 4 || j >= n {
		return
	}

	i++
	num := 0

	//只选一个
	num = num*10 + int(s[j]-'0')
	curr = append(curr, s[j])
	curr = append(curr, '.')
	j++
	helper(curr, i, j, s)
	if j >= n || num == 0 {
		return
	}

	//选两个
	num = num*10 + int(s[j]-'0')
	curr[len(curr)-1] = s[j]
	curr = append(curr, '.')
	j++
	helper(curr, i, j, s)
	if j >= n {
		return
	}

	//选3个
	num = num*10 + int(s[j]-'0')
	if num > 255 {
		return
	}
	curr[len(curr)-1] = s[j]
	curr = append(curr, '.')
	j++
	helper(curr, i, j, s)
}
