package array

/*
杨辉三角
给定一个非负数num，生成杨辉三角的前num行。数据范围：1≤num≤30
                      1
                   1     1
 				1     2     1
             1     3     3     1
*/

//YangHuiTriangle 利用动态规划生成前num行杨辉三角
func YangHuiTriangle(num int) [][]int {
	res := make([][]int, num+1)
	res[0], res[1] = []int{}, []int{1}

	for i := 2; i <= num; i++ {
		res[i] = make([]int, i)
		res[i][0] = 1
		for j := 1; j < i-1; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
		res[i][i-1] = 1
	}

	return res[1:]
}
