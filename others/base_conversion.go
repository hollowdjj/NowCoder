package others

/*
进制转换
给定一个十进制数 M ，以及需要转换的进制数 N 。将十进制数 M 转化为 N 进制数。
当 N 大于 10 以后， 应在结果中使用大写字母表示大于 10 的一位，如 'A' 表示此位为 10 ， 'B' 表示此位为 11 。
若 M 为负数，应在结果中保留负号。

数据范围：2≤N≤16
*/

func BaseConversion(M int, N int) string {
	//1. 已知某4位N进制数a4a3a2a1，它的十进制数值 = a4*N^3 + a3*N^2 + a2*N^1 + a1*N^0
	//2. 已知某十进制数M时，将其转换为N进制数就是上面这个公式的逆过程。即M%N为最低位，(M/N)%N为次低位......
	sign := "+"
	if M < 0 {
		sign = "-"
		M = -M
	}
	flag := "0123456789ABCDEF"
	res := ""
	for M > 0 {
		res = string(flag[M%N]) + res
		M /= N
	}
	if sign == "-" {
		return sign + res
	}
	return res
}
