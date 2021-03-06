package others

/*
最大公约数

如果有一个自然数a能被自然数b整除，则称a为b的倍数， b为a的约数。几个自然数公有的约数，叫做这几个自然数的公约数。
公约数中最大的一个公约数，称为这几个自然数的最大公约数。

输入a和b , 请返回a和b的最大公约数。

进阶：空间复杂度O(1)，时间复杂度O(logn)
*/

func Gcd(a, b int) int {
	//这道题肯定是不能用暴力法的，会超时。下面使用的是一个被称为辗转相除法的算法，以434和652为例：
	// 1. 434 / 652 = 0(余434)
	// 2. 652 / 434 = 1(余218)
	// 3. 434 / 218 = 1(余216)
	// 4. 218 / 216 = 1(余2)
	// 5. 216 / 2 = 108(余0)
	//当余数为0时，当前除数即为最大公约数。
	if a%b == 0 {
		return b
	}
	return Gcd(b, a%b)
}
