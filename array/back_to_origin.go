package array

/*
回到原点

圆环上有10个点，编号为0~9。从0点出发，每次可以逆时针和顺时针走一步，问走n步回到0点共有多少种走法。
输入: 2
输出: 2
解释：有2种方案。分别是0->1->0和0->9->0
*/

func BackToOrigin(n int) int {
	//走n步到0的方案数=走n-1步到1的方案数+走n-1步到9的方案数。
	//dp[i][j]表示走i步到编号为j的点的方案数。那么状态转移方程为：
	//dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
	length := 10
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= 9; j++ {
			dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
		}
	}
	return dp[n][0]
}
