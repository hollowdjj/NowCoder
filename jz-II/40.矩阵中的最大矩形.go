package jz_II

/*
输入：matrix = ["10100","10111","11111","10010"]
输出：6
解释：最大矩形如上图所示。
*/

func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 {
		return 0
	}
	n, m := len(matrix), len(matrix[0])
	grid := make([][]int, n)

	for i := range matrix {
		grid[i] = make([]int, m)
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				continue
			}
			if j == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i][j-1] + 1
			}
		}
	}
	//fmt.Println(grid)
	//对每一列计算最大矩形面积
	res := 0
	for j := 0; j < m; j++ {
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = grid[i][j]
		}
		//fmt.Println(nums,maxArea(nums))
		res = max(res, maxArea(nums))
	}
	return res
}

func maxArea(heights []int) int {
	i, n := 0, len(heights)
	var stack []int

	res := 0
	for i < n {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := 0
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			res = max(res, width*heights[top])
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
