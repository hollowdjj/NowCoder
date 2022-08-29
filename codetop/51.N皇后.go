package codetop

/*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数n，返回所有不同的n皇后问题的解决方案。
每一种解法包含一个不同的n皇后问题 的棋子放置方案，该方案中'Q'和'.'分别代表了皇后和空位。

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
*/

func solveNQueens(n int) [][]string {
	//pos[i]表示的是第i行的pos[i]列放置皇后
	pos := make([]int, n)
	var res [][]string
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, getRes(pos))
			return
		}
		for column := 0; column < n; column++ {
			if valid(row, column, pos) {
				pos[row] = column //放置皇后
				backtrack(row + 1)
			}
		}
	}
	backtrack(0)
	return res
}

func getRes(pos []int) []string {
	n := len(pos)
	res := make([]string, n)
	for i := range pos {
		temp := make([]byte, n)
		for j := 0; j < n; j++ {
			if pos[i] == j {
				temp[j] = 'Q'
			} else {
				temp[j] += '.'
			}
		}
		res[i] = string(temp)
	}
	return res
}

//判断[row,column]位置处能否放置皇后
func valid(row, column int, pos []int) bool {
	//因为是从上到下摆放皇后的，故只需关注同一列，左上斜线，右上斜线上是否有皇后
	upLeft, upRight := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if pos[i] == column {
			return false
		}
		if upLeft >= 0 && pos[i] == upLeft {
			return false
		}
		if upRight < len(pos) && pos[i] == upRight {
			return false
		}
		upLeft--
		upRight++
	}
	return true
}
