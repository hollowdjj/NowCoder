package string

/*
编辑距离(一)
给定两个字符串 str1 和 str2 ，请你算出将 str1 转为 str2 的最少操作数。
你可以对字符串进行3种操作：
1.插入一个字符
2.删除一个字符
3.修改一个字符。

字符串长度满足1≤n≤500，保证字符串中只出现小写英文字母。

输入："nowcoder","new"
返回值：6
说明："nowcoder"=>"newcoder"(将'o'替换为'e')，修改操作1次
      "nowcoder"=>"new"(删除"coder")，删除操作5次
*/

func EditDistance(str1, str2 string) int {
	//采用动态规划。dp[i][j]表示将s1[0:i]转换成s2[0:j]的最少操作数
	//当s1[i-1] = s2[j-1]时，dp[i][j] = dp[i-1][j-1]。否则需要比较
	//一下三种操作哪一个最少：
	//1. 假设执行插入操作。要求s1[0:i] = s2[0:j-1]，故dp[i][j] = dp[i][j-1]+1
	//2. 假设执行删除操作。要求s1[0:i-1] = s2[0:j]，故dp[i][j] = dp[i-1][j]+1
	//3. 假设执行修改操作。要求s1[0:i-1] = s2[0:j-1]，故dp[i][j] = dp[i-1][j-1]+1
	//故dp[i][j] = Min{dp[i][j-1],dp[i-1][j],dp[i-1][j-1]}+1
	//需要注意的是，dp[i]的长度必须是len(str1)+1，否则不能处理str1的长度为1的情况

	n1, n2 := len(str1), len(str2)
	if n1 == 0 {
		return n2
	}
	dp := make([][]int, n1+1)
	dp[0] = make([]int, n2+1)
	for i := 1; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		//str1[0:i]变成str2[0:0]，也即空字符串的最少操作数就是删除i次
		dp[i][0] = dp[i-1][0] + 1
		for j := 1; j <= n2; j++ {
			//str1[0:0]即空字符串变成str2[0:j]的最少操作就是增加j次
			dp[0][j] = dp[0][j-1] + 1
		}
	}

	//选择
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1])) + 1
			}
		}
	}

	return dp[n1][n2]
}
