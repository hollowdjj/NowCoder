package stack

import (
	"nowcoder/array"
	"nowcoder/utility"
)

/*
直方图最大矩形面积问题，即给定一个直方图，找出该直方图中的最大面积。
例如，输入直方图的高度矩阵[1,4,5,2]，输出8(2*4)

       | |
     | | |
	 | | |
     | | | |
   | | | | |

*/

func MaxArea(heights []int) int {
	//最基础的做法显然就是，首先固定一个矩形，然后分别朝左右两边扩展，直到左右两边都扩展到了某一个高度低于
	//固定的矩形的矩形。 在每一次的扩展当中，都计算一次矩形面积并更新最大面积。
	n := len(heights)
	maxArea := 0
	for i, _ := range heights {
		f1, f2 := false, false
		left, right := i-1, i+1
		for {
			if left < 0 || heights[left] < heights[i] {
				left = left + 1
				f1 = true
			}
			if right > n-1 || heights[right] < heights[i] {
				right = right - 1
				f2 = true
			}
			if f1 && f2 {
				break
			}
			maxArea = array.MaxOfTwoInt(maxArea, heights[i]*(right-left+1))
			left--
			right++
		}
	}
	return maxArea
}

func MaxAreaMonotoneStack(heights []int) int {
	//对每一个矩形，我们都要找他左边以及右边第一个高度小于它的矩形。如果，我们能快速
	//地找到左边第一个高度小于它的矩形，那么就能够节约很多时间。能实现这个功能的数据
	//结构叫单调栈。单调栈分为单调递增栈以及单调递减栈，前者表示从栈底到栈顶，元素大
	//小递增的栈，后者表示从栈底到栈顶，元素大小递减的栈。
	//在这里，我们需要找到第i个矩形左边第一个高度小于它的矩形，因此使用一个单调递增栈
	maxArea := 0
	var stack utility.Stack
	i := 0
	for i < len(heights) {
		//如果栈为空时，i进栈。如果栈不为空，但是栈顶元素表示的矩形高度小于heights[i]，那么i也入栈
		if stack.Empty() || heights[(*stack.Top()).(int)] <= heights[i] {
			stack.Push(i)
			i++
		} else {
			//栈顶元素对应的高度大于heights[i]，说明我们找到了栈顶元素右边第一个小于它的元素。
			//而栈顶元素左边第一个小于它的元素就在它的下方。
			top := (*stack.Pop()).(int)
			width := 0
			if stack.Empty() {
				//如果栈为空，说明top的左边没有比它高度更低的矩形了，此时的宽度就是i
				width = i
			} else {
				//如果栈不为空，那么左边界就是当前栈顶元素
				width = i - (*stack.Top()).(int) - 1
			}
			maxArea = max(maxArea, heights[top]*width)
		}
	}

	//如果栈不是空的，那么依次弹出栈顶元素，并计算面积。这主要是因为，高度数组在最后一部分
	//出现了连续递增的子数组。此时i的值固定为len(heights)-1
	for !stack.Empty() {
		top := (*stack.Pop()).(int)
		width := 0
		if stack.Empty() {
			//如果栈为空，说明top的左边没有比它高度更低的矩形了，此时的宽度就是i
			width = i
		} else {
			//如果栈不为空，那么左边界就是当前栈顶元素
			width = i - (*stack.Top()).(int) - 1
		}
		maxArea = max(maxArea, heights[top]*width)
	}
	return maxArea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
