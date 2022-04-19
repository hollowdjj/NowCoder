package jz

/*
输入一个整数 n ，求 1～n 这 n 个整数的十进制表示中 1 出现的次数
例如， 1~13 中包含 1 的数字有 1 、 10 、 11 、 12 、 13 因此共出现 6 次
*/

func NumberOf1Between1AndN(n int) int {
	//分一下几种情况讨论：
	//1. 当curr位为0时，例如514 0 36。curr指向0，当前位为百位，base=100，a = 514，b = 36。此时，我们向前
	//借一位，也就是向a借一位，并让curr变成1。由于curr只是变成了1，还可以提供后面的数的借位，因此b的取值范围
	//从[0,36]变成了[0,99]，a的范围变成了[0,513]。因此1的个数为：a * base
	//2. 当curr位为1时，例如514 1 36。我们可以选择借位或者不借位。因此，1的个数为: 1*(b+1) + a*base
	//3. 当curr位大于1时，例如514 4 36。当前位可以提供后面的借位。因此，1的个数为：(a+1)*base

	res := 0
	for base := 1; base <= n; base *= 10 {
		curr := n / base % 10
		a, b := n/base/10, n%base
		if curr == 0 {
			res += a * base
		} else if curr == 1 {
			res += b + 1
			res += a * base
		} else {
			res += (a + 1) * base
		}
	}
	return res
}
