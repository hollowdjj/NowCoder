package stack

import (
	"nowcoder/utility"
)

/*
给定一个仅包含 0 和 1 ，大小为 n*m 的二维二进制矩阵，找出仅包含 1 的最大矩形并返回面积。

数据范围： 保证输入的矩形中仅含有 0 和 1

例如输入[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]时，对应的输出为8，所形成的最大矩形如下图所示：

1  0  1  0  0
1  1  1  1  0
1  1  1  1  0
1  0  0  1  0
*/

func MaximalRectangle(matrix [][]int) int {
	//思路在于，使用一个辅助矩阵a。其中a[i][j]表示matrix[i][j]左边，包括它自己在内连续1的个数
	//进而形成了一个个直方图。然后对矩阵a的每一列求一个直方图的最大矩形面积然后得出最大面积。该
	//最大面积即为解。

	//生成辅助矩阵a
	n, m := len(matrix), len(matrix[0])
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		if matrix[i][m-1] == 1 {
			a[i][m-1] = 1
		}
		for j := m - 2; j >= 0; j-- {
			if matrix[i][j] == 1 {
				a[i][j] = a[i][j+1] + 1
			}
		}
	}

	//对矩阵的每一列，求直方图的最大矩形面积
	res := 0
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			heights[j] = a[j][i]
		}
		res = max(res, maxRectArea(heights))
	}
	return res
}

//maxArea 求直方图的最大矩形面积
func maxRectArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	} else if n == 1 {
		return heights[0]
	}

	var s utility.Stack //一个单调递增栈
	res := 0
	width := 0
	i := 0
	for i < n {
		if s.Empty() || heights[(*s.Top()).(int)] < heights[i] {
			s.Push(i)
			//只在这里i++的原因在于，对top和top下面的元素而言，其右边都有可能是i。
			i++
		} else {
			//如果heights[top] > heights[i]。那么对索引为top的矩形而言，其左边第一个高度小于
			//它的矩形的索引为top下面那个；其右边第一个高度小于它的矩形的索引为i。此时，就可以计算
			//并更新最大矩形面积
			top := (*s.Pop()).(int)
			if s.Empty() {
				//s为空，说明top左边没有高度小于heights[top]的矩形。那么此时宽度为i，其实就是i-1-0+1
				//即左边可以扩展到0，也即左边界可以认为是-1。那么有i - (-1) -1 = i
				width = i
			} else {
				//栈不为空，那么左边界就是当前栈顶元素。
				width = i - (*s.Top()).(int) - 1
			}
			res = max(res, heights[top]*width)
		}
	}

	//如果s不为空，说明高度数组在最后一段出现了高度连续递增的子序列。依次弹出栈顶元素，计算并更新最大面积即可。
	for !s.Empty() {
		top := (*s.Pop()).(int)
		if s.Empty() {
			width = i
		} else {
			//栈不为空，那么左边界就是当前栈顶元素。
			width = i - (*s.Top()).(int) - 1
		}
		res = max(res, heights[top]*width)
	}
	return res
}
