package jz

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶(n为正整数)总共有多少种跳法。
*/

func jumpFloorII(number int) int {
	//dp[i] = Sum{dp[j]}  j<i。状态压缩后，用prevSum代替Sum{dp[j]}
	prevSum := 2
	curr := 1
	for i := 2; i <= number; i++ {
		curr = prevSum
		prevSum += curr
	}
	return curr
}
