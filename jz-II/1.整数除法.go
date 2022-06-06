package jz_II

import "math"

/*
给定两个整数a和b，求它们的除法的商a/b，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8以及truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,2^31−1]。本题中，如果除法结果溢出，则返回 2^31 − 1
*/

func divide1(a int, b int) int {
	//最简单的方式是通过加法模拟除法，但会超时。
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	res := 0
	//判断符号
	sign := 1
	if a > 0 && b > 0 || a < 0 && b < 0 {
		sign = 1
	} else {
		sign = -1
	}
	a, b = abs(a), abs(b)
	t := b
	for b <= a {
		res++
		b += t
	}

	if sign == -1 {
		return -res
	}
	return res
}

func divide2(a, b int) int {
	//采用二分法
	MIN, MAX := math.MinInt32, math.MaxInt32
	if a == MIN && b == -1 {
		return MAX
	}
	//判断符号
	sign := 1
	if a > 0 && b > 0 || a < 0 && b < 0 {
		sign = 1
	} else {
		sign = -1
	}
	a, b = abs(a), abs(b)
	res := 0
	for b <= a {
		t, c := b, 1
		for t < MAX>>1 && t+t <= a {
			t += t
			c += c
		}
		//例如15/2跳出上一个循环时，t=8 c=4，此时15-8=7还可以继续除2
		a -= t
		res += c
	}
	if sign == -1 {
		return -res
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
