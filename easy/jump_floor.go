package easy

import "fmt"

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
要求：时间复杂度：O(n)，空间复杂度：O(1)
*/

//JumpFloor 动态规划
func JumpFloor(number int) int {
	//初始化dp table
	dp := make([]int, number+1)
	dp[0] = 1
	dp[1] = 1
	//选择
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

func TestJumpFloor(number int) {
	fmt.Println(JumpFloor(number))
}
