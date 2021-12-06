package array

/*
螺旋矩阵
给定一个mxn大小的矩阵，按螺旋顺序返回矩阵中的所有元素。
要求：空间复杂度O(mn)，时间复杂度O(mn)
例如：
		1  2  3
        4  5  6    ------>   1,2,3,6,9,8,7,4,5
		7  8  9
*/

//SpiralOrder 螺旋打印矩阵
func SpiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	res := make([]int, m*n)

	top, bottom, left, right := 0, m-1, 0, n-1
	k := 0
	for k < m*n {
		//上面，从左到右打印
		for i := left; i <= right; i++ {
			res[k] = matrix[top][i]
			k++
		}
		top++
		//右边，从上到下打印
		for i := top; i <= bottom; i++ {
			res[k] = matrix[i][right]
			k++
		}
		right--
		//下边，从右往左打印。top<bottom是为了免只有一行时的特殊处理
		for i := right; top <= bottom && i >= left; i-- {
			res[k] = matrix[bottom][i]
			k++
		}
		bottom--
		//左边，从下往上打印。left < right则是只有一列时的特殊处理
		for i := bottom; left <= right && i >= top; i-- {
			res[k] = matrix[i][left]
			k++
		}
		left++
	}
	return res
}
