package string

import "strconv"

/*
最大值
有一个只由字符'1'到'9'组成的长度为n的字符串 ss ，现在可以截取其中一段长度为k
的子串并且将该子串当作十进制的正整数，如对于子串"123"，其对应的十进制数字就是123123 。

如果想让这个正整数尽可能的大的话，问该正整数最大能是多少。
函数传入一个长度为n的字符串s和一个正整数k，请你返回答案。

数据范围：1≤n≤10^5,1≤k≤min(n,8)
例如："321",2 返回值：32
说明：所有长度为 22 的子串为："32"和"21"，显然3232是最大的。
*/

func MaxValue(s string, k int) int {
	//采用滑动窗口。因为题目限制了k的最大值为8.因此，可以使用strconv
	res := -1
	n := len(s)
	for i := 0; i <= n-k; i++ {
		num, _ := strconv.Atoi(s[i : i+k])
		if num > res {
			res = num
		}
	}
	return res
}
