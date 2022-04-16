package jz

/*
斐波那契数列：f(x) = f(x-1)+f(x-2) x > 2;    f(x) = 1 x=1,2
*/

func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	prev, prevOfPrev := 1, 1
	current := 0
	for x := 3; x <= n; x++ {
		current = prev + prevOfPrev
		prevOfPrev = prev
		prev = current
	}
	return current
}
