package string

/*
编辑距离(二)
给定两个字符串str1和str2，再给定三个整数ic，dc和rc，分别代表插入、删除和替换一个字符的代价，请输出将str1编辑成str2的最小代价。

数据范围：0 ≤∣str1∣,∣str2∣≤5000，0≤ic,dc,rc≤10000
要求：空间复杂度O(n)，时间复杂度O(nlogn)
*/

func MinEditCost(str1 string, str2 string, ic int, dc int, rc int) int {
	//dp[i][j]表示将字符串str1[0:i]编辑为str2[0:j]的最小代价。故，状态转移方程为：
	//当str1[i-1] = str2[j-1]时，dp[i][j] = dp[i-1][j-1]
	//而当str1[i-1] != str2[j-1]时，需要分三种情况：
	//1. 若str1执行插入操作。要求str1[0:i] = str2[0:j-1]，故dp[i][j] = dp[i][j-1] + ic
	//2. 若str1执行删除操作。要求str1[0:i-1] = str2[0:j]，故dp[i][j] = dp[i-1][j] + dc
	//3. 若str1执行替换操作。要求str1[0:i-1] = str2[0:j-1]，故dp[i][j] = dp[i-1][j-1]+rc
	//此时dp[i][j] = Min{dp[i][j-1],dp[i-1][j],dp[i-1][j-1]}

	n1, n2 := len(str1), len(str2)
	if n1 == 0 {
		return ic * n2
	}
	//初始化dp数组
	dp := make([][]int, n1+1)
	dp[0] = make([]int, n2+1)
	for i := 1; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][0] = dp[i-1][0] + dc
		for j := 1; j <= n2; j++ {
			dp[0][j] = dp[0][j-1] + ic
		}
	}

	//选择
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+ic, min(dp[i-1][j]+dc, dp[i-1][j-1]+rc))
			}
		}
	}

	return dp[n1][n2]
}
