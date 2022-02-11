package string

import "strings"

/*
旋转字符串
给定两字符串A和B，如果能将A从中间某个位置分割为左右两部分字符串（可以为空串），
并将左边的字符串移动到右边字符串后面组成新的字符串可以变为字符串B时返回true。

例如：如果A=‘youzan’，B=‘zanyou’，A按‘you’‘zan’切割换位后得到‘zanyou’和B相同，返回true。
再如：如果A=‘abcd’，B=‘abcd’，A切成‘abcd’和''（空串），换位后可以得到B，返回true。

数据范围：A,B字符串长度满足n≤1000，保证字符串中仅包含小写英文字母和阿拉伯数字
进阶： 时间复杂度 O(n)，空间复杂度O(n)
*/

func Rotate(A, B string) bool {
	//这道题的解法非常巧妙。只要B是A+A的一个子串时，返回true。
	//例如：youzanyouzan zanyou
	if len(A) != len(B) {
		return false
	}
	return strings.Contains(A+A, B)
}
