package array

/*
给定一个整数数组nums和一个整数target，请你返回该数组能构成多少种不同的表达式等于target。
规则如下：
1.将数组里每个整数前面可以添加"+"或者"-"符号，组成一个表达式，例如[1,2]，可以变成”+1+2","+1-2","-1+2","-1-2"，这四种
2.只能添加"+"与"-"符号，不能添加其他的符号
3.如果构不成等于target的表达式，请返回0
4.保证返回的结果个数在整数范围内

输入：[1,1,1,2],3
输出：3
说明：
-1 + 1 + 1 + 2 = 3
+1 - 1 + 1 + 2 = 3
+1 + 1 - 1 + 2 = 3
*/

func FindTargetSumWays(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := 0
	var dfs func(i int, rest int)
	dfs = func(i int, rest int) {
		if i == n {
			if rest == 0 {
				res += 1
			}
			return
		}
		dfs(i+1, rest-nums[i])
		dfs(i+1, rest+nums[i])
	}
	dfs(0, target)
	return res
}
