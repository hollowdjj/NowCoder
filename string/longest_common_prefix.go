package string

/*
最长公共前缀
给你一个大小为 n 的字符串数组strs ，其中包含n个字符串 , 编写一个函数来查找
字符串数组中的最长公共前缀，返回这个公共前缀。

数据范围： 0≤n≤5000， 0≤len≤5000
*/

func LongestCommonPrefix(strs []string) string {
	//垂直扫描
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	if strs[0] == "" {
		return ""
	}

	b := make([]byte, 0)
	column := 0
	length := len(strs[0])
	prev := strs[0][column]
	for column < length {
		for i, _ := range strs {
			length = min(length, len(strs[i]))
			if strs[i][column] != prev {
				return string(b)
			}
		}
		column++
		b = append(b, prev)
		if column == length {
			return string(b)
		}
		prev = strs[0][column]
	}
	return string(b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
