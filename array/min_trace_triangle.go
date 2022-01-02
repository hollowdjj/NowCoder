package array

/*
给定一个正三角形数组，自顶到底分别有 1，2，3，4，5...，n 个元素，找出自顶向下的最小路径和。
每一步只能移动到下一行的相邻节点上，相邻节点指下行中下标与之相同或下标加一的两个节点。
例如当输入[[2],[3,4],[6,5,7],[4,1,8,3]]时，对应的输出为11，
所选的路径如下图所示：
2
3 4
6 5 7
4 1 8 3
*/

func MinTrace(triangle [][]int) int {
	//dp数组的定义为：dp[i][j]为移动到triangle[i][j]元素的最短路径和
	//从而状态转移方程为：
	//
	//   dp[i][j] = Min{dp[i-1][j],dp[i-1][j-1]; when i-1>=0, j-1>=0} + triangle[i][j]
	n := len(triangle)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = triangle[i][0] + dp[i-1][0]
	}

	//选择
	for i := 1; i < n; i++ {
		for j := 1; j <= i; j++ {
			if j < i {
				//上一层最多只有第i-1列
				dp[i][j] = MinOfTwoInt(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
	}

	//遍历最后一行，找到最小值
	res := dp[n-1][0]
	for i := 1; i < len(triangle); i++ {
		if dp[n-1][i] < res {
			res = dp[n-1][i]
		}
	}
	return res
}
