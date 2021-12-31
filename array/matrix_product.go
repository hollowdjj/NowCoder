package array

/*
给定两个 n*n 的矩阵 A 和 B ，求 A*B 。

要求：空间复杂度O(n^2)，时间复杂度 O(n^3 )
进阶：本题也有空间复杂度O(n^2)，时间复杂度 O(n^{log7})的解法
PS：更优时间复杂度的算法这里并不考察
*/

func MatrixProduct(a [][]int, b [][]int) [][]int {
	//初始化一个n*n的矩阵
	n := len(a)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				//a的第i行元素与b的第j列元素对应相乘
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res
}
