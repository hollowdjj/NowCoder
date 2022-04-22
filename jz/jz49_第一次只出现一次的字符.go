package jz

/*
在一个长为 字符串中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）

输入："google"
输出：4
*/

func FirstNotRepeatingChar(str string) int {
	dic := make(map[rune]int)
	for i, r := range str {
		if _, ok := dic[r]; ok {
			delete(dic, r)
		} else {
			dic[r] = i
		}
	}

	res := len(str) + 1
	for _, v := range dic {
		if v < res {
			res = v
		}
	}
	if res == len(str)+1 {
		return -1
	}
	return res
}
