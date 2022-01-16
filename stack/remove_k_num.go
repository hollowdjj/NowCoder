package stack

import (
	"nowcoder/utility"
	"strconv"
	"strings"
)

/*
给定一个以字符串表示的数字 num 和一个数字 k ，从 num 中移除 k 位数字，使得剩下的数字最小。如果可以删除全部数字则剩下 0

数据范围：num的长度满足1，保证 num 中仅包含 0~9 的十进制数

输入：“1432219”，3
返回值：“1219”
说明：移除4 3 2 后剩下1219

输入： “10”，1
返回值：“0”
*/

func RemoveKNums(num string, k int) string {
	//思路在于要使得高位的元素尽量小。因此，维护一个单调不减栈(即在单调递增栈的基础上允许出现相等的元素)
	//即，当前数字大于栈顶元素时，push当前元素；当前数字小于栈顶元素时，pop栈顶元素，然后push当前元素，并
	//k--
	var s utility.Stack
	for _, c := range num {
		digit := (int)(c - '0')
		for !s.Empty() && k > 0 && (*s.Top()).(int) > digit {
			s.Pop()
			k--
		}
		s.Push(digit)
	}
	//如果删除次数k还没有用完，那么一直pop直到k为0
	for !s.Empty() && k > 0 {
		s.Pop()
		k--
	}
	//组成数字字符串
	var temp utility.Stack
	for !s.Empty() {
		temp.Push((*s.Pop()).(int))
	}
	var build strings.Builder
	for !temp.Empty() {
		build.WriteString(strconv.Itoa((*temp.Pop()).(int)))
	}
	result := build.String()
	//删除前导0
	result = strings.TrimLeft(result, "0")
	if result == "" {
		return "0"
	}
	return result
}
