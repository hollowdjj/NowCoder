package jz_II

/*
给定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1,col1) ，右下角为 (row2,col2) 。

实现 NumMatrix 类：
NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2)返回左上角(row1,col1)
右下角(row2,col2)的子矩阵的元素总和。
*/

type NumMatrix struct {
	sum [][]int //sum[i][j]表示martrix[i][0..j]的和
}

func Constructor1(matrix [][]int) NumMatrix {
	//一维前缀和
	n, m := len(matrix), len(matrix[0])
	res := NumMatrix{sum: make([][]int, n)}
	for i := 0; i < n; i++ {
		res.sum[i] = make([]int, m)
		res.sum[i][0] = matrix[i][0]
		for j := 1; j < m; j++ {
			res.sum[i][j] = res.sum[i][j-1] + matrix[i][j]
		}
	}
	return res
}

func (m *NumMatrix) SumRegion1(row1 int, col1 int, row2 int, col2 int) int {
	//根据前缀和，计算每一行的和然后相加
	res := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			res += m.sum[i][col2]
		} else {
			res += m.sum[i][col2] - m.sum[i][col1-1]
		}
	}
	return res
}

func Constructor2(matrix [][]int) NumMatrix {
	//二维前缀和。求子矩阵的和，可以用分块矩阵的思想。例如：
	//   a1  a2  a3            a5  a6
	//   a4  a5  a6  求子矩阵:   a8  a9 的和。
	//   a7  a8  a9
	//设sum(ai,aj)为分别以ai,sj为左上角和右下角的矩阵元素和。那么有：
	//sum(a5,a9)=sum(a1,a9)-sum(a1,a7)-sum(a1,a3) + sum(a1,a1)
	//因此，使用一个辅助矩阵sum，其中sum(i,j)表示以matrix[0][0]为左上角，
	//matrix[i][j]为右下角的矩阵元素和。
	n, m := len(matrix), len(matrix[0])
	sum := make([][]int, n+1)
	sum[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		sum[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{sum}
}

func (m *NumMatrix) SumRegion2(row1 int, col1 int, row2 int, col2 int) int {
	return m.sum[row2+1][col2+1] - m.sum[row2+1][col1] - m.sum[row1][col2+1] + m.sum[row1][col1]
}
