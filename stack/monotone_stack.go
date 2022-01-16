package stack

import "nowcoder/utility"

func foundMonotoneStackClassic(nums []int) [][]int {
	//单调栈的经典做法是，首先从左往右遍历，得到每个数左边第一个小于它的数
	//然后，从右往左遍历，得到每个数右边第一个小于它的数。
	n := len(nums)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, 2)
	}
	var s utility.Stack
	for i := 0; i < n; i++ {
		for !s.Empty() || nums[(*s.Top()).(int)] >= nums[i] {
			//保证栈顶元素一定是小于nums[i]的
			s.Pop()
		}
		if s.Empty() {
			//如果栈为空，对于nums[i]而言，其左边不存在小于它的数
			res[i][0] = -1
		} else {
			//否则，nums[i]左边第一个小于它的数的索引为s.Top()
			res[i][0] = nums[(*s.Top()).(int)]
		}
		s.Push(i)
	}
	for !s.Empty() {
		s.Pop()
	}
	for i := n - 1; i >= 0; i-- {
		for !s.Empty() || nums[(*s.Top()).(int)] >= nums[i] {
			//保证栈顶元素一定是小于nums[i]的
			s.Pop()
		}
		if s.Empty() {
			//如果栈为空，对于nums[i]而言，其右边不存在小于它的数
			res[i][1] = -1
		} else {
			//否则，nums[i]右边第一个小于它的数的索引为s.Top()
			res[i][1] = nums[(*s.Top()).(int)]
		}
		s.Push(i)
	}
	return res
}

func foundMonotoneStack(nums []int) [][]int {
	//使用一个单调递增的栈
	n := len(nums)
	if n == 1 {
		return [][]int{{-1, -1}}
	}
	var s utility.Stack
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = []int{-1, -1}
	}
	for i := 0; i < n; {
		if s.Empty() || nums[(*s.Top()).(int)] < nums[i] {
			s.Push(i)
			i++
		} else if nums[(*s.Top()).(int)] > nums[i] {
			//对top而言，其右边界为i，左边界为其下方的元素
			top := (*s.Pop()).(int)
			right, left := i, -1
			if !s.Empty() {
				left = (*s.Top()).(int)
			}
			res[top] = []int{left, right}
		} else {
			i++
		}
	}
	for !s.Empty() {
		top := (*s.Pop()).(int)
		left, right := -1, -1
		if !s.Empty() {
			left = (*s.Top()).(int)
		}
		res[top] = []int{left, right}
	}

	return res
}
