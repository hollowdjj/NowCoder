package others

import "math"

/*
反转数字
给定一个32位的有符号整数num，将num中的数字部分反转，最后返回反转的结果
1.只反转数字部分，符号位部分不反转
2.反转后整数num超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，返回 0
3.假设本题不允许存储 64 位整数(有符号或无符号，即C++不能使用long long ，Java不能使用long等)

数据范围:
-2^31 <= x <= 2^31-1
*/
func Reverse(x int) int {
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	res *= flag
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
