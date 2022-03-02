package array

/*
分糖果问题

一群孩子做游戏，现在请你根据游戏得分来发糖果，要求如下：

1. 每个孩子不管得分多少，起码分到一个糖果。
2. 任意两个相邻的孩子之间，得分较多的孩子必须拿多一些糖果。(若相同则无此限制)

给定一个数组arr代表得分数组，请返回最少需要多少糖果。

输入：[1,1,2]
输出：4
说明：最优分配方案为1,1,2
*/

func Candy(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	//首先正序遍历数组。只要arr[i] > arr[i-1]那么dp[i]=dp[i-1]+1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}

	//然后再逆序遍历数组。只要arr[i-1] > arr[i]，那么dp[i-1] = max(dp[i]+1,dp[i-1])
	res := dp[n-1]
	for i := n - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			//例如[1,2,3,4,2,1]正序遍历完后的sum数组为[1,2,3,4,1,1]
			//那么4=max(4,2+1)
			dp[i-1] = max(dp[i-1], dp[i]+1)
		}
		res += dp[i-1]
	}

	return res
}
