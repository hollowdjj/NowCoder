package jz_II

/*
请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：
要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
*/

func dailyTemperatures(temperatures []int) []int {
	//单调栈
	res := make([]int, len(temperatures))
	var s []int
	for i, t := range temperatures {
		for len(s) > 0 && temperatures[s[len(s)-1]] < t {
			res[s[len(s)-1]] = i - s[len(s)-1]
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return res
}
