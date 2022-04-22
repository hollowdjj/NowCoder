package jz

/*
有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。
我们把一个字符串编码成一串数字，再考虑逆向编译成字符串。
由于没有分隔符，数字编码成字母可能有多种编译结果，例如 11 既可以看做是两个 'a' 也可以看做是一个 'k' 。
但 10 只可能是 'j' ，因为 0 不能编译成任何结果。现在给一串数字，返回有多少种可能的译码结果

输入："12"
输出：2
说明：2种可能的译码结果（”ab” 或”l”）
*/

func EncodedNumToString(nums string) int {
	//dp[i]表示nums[0..i]译码结果数量
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		num := int(nums[i-1]-'0')*10 + int(nums[i]-'0')
		if nums[i] == 0 && num > 26 || num == 0 {
			//诸如30 40...以及00的情况
			return 0
		}
		if nums[i] != '0' {
			dp[i] = dp[i-1]
		}
		if num >= 10 && num <= 26 {
			if i == 1 {
				dp[i]++
			} else {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n-1]
}
