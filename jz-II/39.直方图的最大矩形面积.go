package jz_II

func largestRectangleArea(heights []int) int {
	var stack []int
	i, n := 0, len(heights)
	res := 0
	for i < n {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			//此时，对于栈顶元素而言，i为其右边第一个高度大于它的矩阵的索引
			//而其下方的元素就是左边第一个高度小于等于他的元素的索引，所以这里
			//没有i++，也是因为下方元素的高度可能等于栈顶元素的高度。
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := 0
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			res = max(res, heights[top]*width)
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := 0
		if len(stack) == 0 {
			width = n
		} else {
			width = n - stack[len(stack)-1] - 1
		}
		res = max(res, width*heights[top])
	}

	return res
}
