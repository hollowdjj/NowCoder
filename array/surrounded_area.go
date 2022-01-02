package array

/*
给定一个 n*m 大小的的矩阵，矩阵中由 ‘X' 和 'O' 构成，找到所有被 'X' 围绕的区域，并将其用 'X' 填充。

例如：
[['X','X','X','X'],
['X','O','O','X'],
['X','O','X','X'],
['X','X','O','X']]
中间的三个 ‘O’ 被 'X'围绕，因此将其填充为 'X' ，但第四行的 'O' 下方没有被 'X' 围绕，因此不改变，结果为
[['X','X','X','X'],
['X','X','X','X'],
['X','X','X','X'],
['X','X','O','X']]
*/

func SurroundedArea(board [][]byte) [][]byte {
	n, m := len(board), len(board[0])
	//第一行，最后一行，第一列以及最后一列上的O是不变的，因此只需遍历中间的元素
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if board[i][j] == 'O' {
				//该位置的上，下，左，右都必须出现X，否则不变
				flag := false
				//上
				for k := 0; k < i; k++ {
					if board[k][j] == 'X' {
						flag = true
						break
					}
				}
				if flag == false {
					continue
				}
				flag = false
				//下
				for k := i + 1; k < n; k++ {
					if board[k][j] == 'X' {
						flag = true
						break
					}
				}
				if flag == false {
					continue
				}
				flag = false
				//左
				for k := 0; k < j; k++ {
					if board[i][k] == 'X' {
						flag = true
						break
					}
				}
				if flag == false {
					continue
				}
				flag = false
				//右
				for k := j + 1; k < m; k++ {
					if board[i][k] == 'X' {
						flag = true
						break
					}
				}
				if flag == true {
					board[i][j] = 'X'
				}
			}
		}
	}

	return board
}
