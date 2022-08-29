package jz

/*
实现函数 double Power(double base, int exponent)，求base的exponent次方。

注意：
1.保证base和exponent不同时为0。
2.不得使用库函数，同时不需要考虑大数问题
3.有特殊判题，不用考虑小数点后面0的位数。

例如：
输入：2.10000,3
输出：9.26100
*/

func Power(b float64, n int) float64 {
	//使用非递归的快速幂。当n=6时，n的二进制表示为110。那么有：
	//6 = 0*2^0 + 1*2^1 + 1*2^2 ----> x^6 = x^(0*2^0 + 1*2^1 + 1*2^2)，化简可得：
	//x^6 = x^(0*2^0) * x^(1*2^1) * x^(1*2^2)。也就是说，我们只需要遍历n的二进制数，
	//当第i位数为1时，就将最后的答案乘以x^(2^i)
	if n < 0 {
		b = 1 / b
		n = -n
	}

	var res float64 = 1
	for n > 0 {
		if n&1 == 1 {
			res *= b
		}
		b *= b //x^(2^i) = ((x^2)^2)^2
		n >>= 1
	}
	return res
}
