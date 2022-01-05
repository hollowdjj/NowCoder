package array

/*
给定一个大小为 n*m 的矩阵，请以对角线遍历并返回遍历结果

1  2  3
4  5  6   --->  1,2,4,7,5,3,6,8,9
7  8  9

1  2  3  4
5  6  7  8      1,2,5,6,3,7,7,8
*/

func DiagonalOrder(mat [][]int) []int {
	n, m := len(mat), len(mat[0])
	res := make([]int, n*m)

	index := 0
	i, j := 0, 0
	for index < n*m {
		res[index] = mat[i][j]
		index++
		//如果i+j为偶数的话，遍历方向是从下到上。是奇数的话，遍历方向是从上到下
		//但是，这里需要注意的是，当从下到上遍历时，if j == m-1要在 if i==0判断
		//条件的上面。因为在遍历主对角线时，最后是i==0 && j == m-1，此时要往下走。
		//同样的，当从上到下遍历时，if i == n-1要在if j == 0之前
		if (i+j)%2 == 0 {
			if j == m-1 {
				//如果走到了最右边，往下
				i++
			} else if i == 0 {
				//如果走到了最上边，往右
				j++
			} else {
				//从下到上对角线遍历
				i--
				j++
			}
		} else {
			if i == n-1 {
				//如果走到了最下方，那么就往右
				j++
			} else if j == 0 {
				//如果走到了最左边，那么就往下
				i++
			} else {
				//从上到下对角线遍历
				i++
				j--
			}
		}
	}
	return res
}
