package jz

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，例如，如果输入如下4 X 4矩阵：
[[1,2,3,4],
[5,6,7,8],
[9,10,11,12],
[13,14,15,16]]

[1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
*/

func printMatrix(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, m-1, 0, n-1
	i := 0
	res := make([]int, n*m)
	for i < n*m {
		//从左往右
		for j := left; j <= right && i < n*m; j++ {
			res[i] = matrix[top][j]
			i++
		}
		top++
		//从上到下
		for j := top; j <= bottom && i < n*m; j++ {
			res[i] = matrix[j][right]
			i++
		}
		right--
		//从右到左
		for j := right; j >= left && i < n*m; j-- {
			res[i] = matrix[bottom][j]
			i++
		}
		bottom--
		//从下到上
		for j := bottom; j >= top && i < n*m; j-- {
			res[i] = matrix[j][left]
			i++
		}
		left++
	}
	return res
}
