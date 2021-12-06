package array

/*
斐波那契数列
现在要求输入一个正整数n，输出斐波那契数列的第n项
要求：时间复杂度O(n)，空间复杂度O(1)
*/

//Fibonacci 斐波拉契数列
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		b, a = a+b, b
	}
	return b
}
