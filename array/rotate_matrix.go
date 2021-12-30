package array

/*
顺时针旋转矩阵
有一个N x N的矩阵，请编写一个算法将矩阵顺时针旋转90度
要求：时间复杂度O(n^2)，空间复杂度O(n^2)
进阶：时间复杂度O(n^2)，空间复杂度O(1)
*/

func RotateMatrix(mat [][]int, n int) [][]int {
	//使用一个辅助矩阵，故空间复杂度为O(n^2)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	//遍历所有元素，故时间复杂度为O(n^2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//res中的第i行元素为mat中的第i列元素的倒序
			res[i][j] = mat[n-1-j][i]
		}
	}
	return res
}

func RotateMatrixAdvanced(mat [][]int, n int) [][]int {
	/*
		先沿着对角线交换:
		1  2  3           1  4  7
		4  5  6    ---->  2  5  8
		7  8  9			  3  6  9
	*/
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = mat[j][i]
		}
	}
	//然后再把每一行的数字逆转
	//for i := 0; i < n; i++ {
	//	left, right := 0, n-1
	//	for left < right {
	//		temp := mat[i][left]
	//		mat[i][left] = mat[i][right]
	//		mat[i][right] = temp
	//		left++
	//		right--
	//	}
	//}

	return mat
}
