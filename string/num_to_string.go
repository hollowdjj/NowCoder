package string

/*
把数字翻译成字符串

有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。
现在给一串数字，返回有多少种可能的译码结果
*/

func DecodeNumToString(nums string) int {
	//dp[i]表示nums[0...i]的译码结果数量。状态转移方程为：
	n := len(nums)
	if n == 1 {
		return int(nums[0] - '0')
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		//只要当前数字不是0，那么至少就有dp[i-1]种结果。但是我们还要考虑类似12这种情况
		if nums[i] != '0' {
			dp[i] = dp[i-1]
		}

		//与前一个字符组合起来的数字
		num := (nums[i-1]-'0')*10 + nums[i] - '0'

		//11102 当前数字为0时，dp[i] = dp[i-2]
		if num >= 10 && num <= 26 {
			if i == 1 {
				dp[i]++
			} else {
				dp[i] += dp[i-2]
			}
		} else if nums[i] == '0' && (num > 26 || num == 0) {
			//出现了00 30 40 50 ....这样的子串时，没有译码结果
			return 0
		}
	}

	return dp[n-1]
}
