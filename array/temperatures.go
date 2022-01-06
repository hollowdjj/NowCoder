package array

import "nowcoder/utility"

/*
根据往后 n 天的天气预报，计算每一天需要等待几天才会出现一次更高的气温，如果往后都没有更高的气温，则用 0 补位。

例如往后三天的气温是 [1,2,3] ， 则输出 [1,1,0]
*/

func Temperatures(temperatures []int) []int {
	//这道题考察的是单调栈的知识。具体来说，就是维护一个温度值下标的单调栈，使得从
	//栈顶到栈底的元素对应的温度值依次递减。对于温度列表中的所有元素，t[i]。如果栈
	//为空，那么元素i进栈；如果栈不为空，那么比较栈顶元素top对应的温度值t[top]与
	//t[i]的大小:
	//1. 如果t[i] > t[top]，那么top出栈，且top对应的等待天数为i-top。此时，我们需要
	//比较栈顶元素对应的温度值与t[i]的大小，直到t[i]<t[top]或栈为空，然而i进栈。
	//2. 如果t[i] < t[top]，那么i入栈。
	res := make([]int, len(temperatures))
	var stack utility.Stack
	for i, _ := range temperatures {
		if stack.Empty() || temperatures[i] < (*stack.Top()).(int) {
			stack.Push(i)
		} else {
			//一直pop直到t[i] < t[top]或者栈为空
			for {
				if stack.Empty() || temperatures[i] < temperatures[(*stack.Top()).(int)] {
					stack.Push(i)
					break
				}
				top := (*stack.Top()).(int)
				res[top] = i - top
				stack.Pop()
			}
		}
	}
	return res
}
