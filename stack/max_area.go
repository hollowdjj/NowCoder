package stack

import "nowcoder/array"

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
		left, right := 0, 0
		if i > 0 {
			left = i - 1
		}
		if i < n-1 {
			right = i + 1
		}
		for left >= 0 && right < n {
			if heights[left] >= heights[i] && heights[right] >= heights[i] {
				maxArea = array.MaxOfTwoInt(maxArea, heights[i]*(right-left))
				left--
				right++
			} else if heights[left] >= heights[i] && heights[right] < heights[i] {
				maxArea = array.MaxOfTwoInt(maxArea, heights[i]*(right-left))
				break
			}
		}
		for left >= 0 {
			maxArea = array.MaxOfTwoInt(maxArea, heights[i]*(right-left))
			left--
		}
		for right < n {
			maxArea = array.MaxOfTwoInt(maxArea, heights[i]*(right-left))
			right++
		}
	}
	return maxArea
}
