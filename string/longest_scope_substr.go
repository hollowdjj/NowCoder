package string

/*
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

例1: 对于字符串 "(()" 来说，最长的格式正确的子串是 "()" ，长度为 2 .
例2：对于字符串 ")()())" , 来说, 最长的格式正确的子串是 "()()" ，长度为 4 .

字符串长度：0≤n≤5∗10^5
要求时间复杂度O(n),空间复杂度O(n)
*/
func LongestValidParentheses(s string) int {
	//dp[i]表示s[0:i]的最长格式正确括号子串长度。状态转移方程：
	//当s[i] = '('时dp[i] = 0(因为必须要是连续子串)
	//当s[i] = ')' && s[i-1] = '('时 dp[i] = dp[i-2] + 2。例如：.....()
	//当s[i] = ')' && s[i-1] == ')时：
	// 形如  ...(      (....)     )
	//     i-1-dp[i-1]     i-1    i
	//如果s[i-1-dp[i-1]]为'('的话dp[i] = dp[i-1]+dp[i-1-dp[i-1]-1]+2
	n := len(s)
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			if i-1 >= 0 && s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}

			} else if i-1 >= 0 && s[i-1] == ')' {
				if dp[i-1] > 0 && i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
					if i-1-dp[i-1]-1 >= 0 {
						dp[i] = dp[i-1] + dp[i-1-dp[i-1]-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func longestValidParentheses(s string) int {
	//在此方法中，我们利用两个计数器left和right 。首先，我们从左到右遍历字符串，对于遇到的每个‘(’，我们
	//增加left计数器，对于遇到的每个‘)’，我们增加right计数器。每当left计数器与right计数器相等时，我们计算
	//当前有效字符串的长度，并且记录目前为止找到的最长子字符串。当right计数器比left 计数器大时，我们将left
	//和right计数器同时变回0。

	//这样的做法贪心地考虑了以当前字符下标结尾的有效括号长度，每次当右括号数量多于左括号数量的时候之前的字符
	//我们都扔掉不再考虑，重新从下一个字符开始计算，但这样会漏掉一种情况，就是遍历的时候左括号的数量始终大于
	//右括号的数量，即 (() ，这种时候最长有效括号是求不出来的。
	//解决的方法也很简单，我们只需要从右往左遍历用类似的方法计算即可，只是这个时候判断条件反了过来：

	//当left计数器比right计数器大时，我们将left和right计数器同时变回0
	//当left计数器与right计数器相等时，我们计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串
	//这样我们就能涵盖所有情况从而求解出答案。

	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func longestValidParentheses1(s string) int {
	//由于是求最长有效括号子串的长度，栈中需要保存当前遍历到的子串中
	//最后一个没有左括号与之匹配的右括号。遇到"("就将其下标压栈，遇
	//到")"括号时，若栈为空，说明这就是当前最后一个没有左括号与之匹配
	//的右括号，将其下标压栈。若栈不为空，那么弹出栈顶元素并计算最大长度
	stack := []int{-1} //())()
	res := 0
	for i, b := range s {
		if b == '(' {
			//遇到左括号就压栈
			stack = append(stack, i)
		} else {
			//弹出栈顶元素
			stack = stack[:len(stack)-1]
			//()()弹出第二个左括号，那么长度就是3-1=2
			if len(stack) > 0 {
				res = max(res, i-stack[len(stack)-1])
			} else {
				//())()
				stack = append(stack, i)
			}
		}
	}
	return res
}
