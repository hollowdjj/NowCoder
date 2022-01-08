package array

/*
在 4x4 的棋盘上摆满了黑白棋子，黑白两色棋子的位置和数目随机，其中0代表白色，1代表黑色；左上角坐标为 (1,1) ，右下角坐标为 (4,4) 。
现在依次有一些翻转操作，要对以给定翻转坐标(x,y)（也即第x行第y列）为中心的上下左右四个棋子的颜色进行翻转。
给定两个数组 A 和 f ，分别代表 初始棋盘 和 哪些要进行翻转的位置(x,y) ，请返回经过所有翻转操作后的棋盘。

例如输入[[0,0,1,1],[1,0,1,0],[0,1,1,0],[0,0,1,0]],[[2,2],[3,3],[4,4]]时，初始键盘如下图所示：

0  0  1  1 			  0  1  1  1		0  1  1  1         0  1  1  1
1  0  1  0 			  0  0  1  0        0  0  1  0		   0  0  1  0
0  1  1  0      ----> 0  1  1  0  ----> 0  1  1  1  ---->  0  1  1  0
0  0  1  0 			  0  0  1  0        0  0  0  0		   0  0  1  0

注意，是翻转，而不是上下互换 左右互换
*/

func FlipChess(A [][]int, f [][]int) [][]int {
	//以A[i][j]为中心，上下左右四个棋子翻转，其实就是swap(A[i][j-1],A[i-1][j])以及swap(A[i+1][j],A[i][j+1])
	//只是需要注意一下,j-1 i-1 i+1 j+1是否存在。这样做的话，那么时间复杂度为O(n)，空间复杂度为O(1)。
	n := len(f)
	for k := 0; k < n; k++ {
		i, j := f[k][0]-1, f[k][1]-1
		if j >= 1 && i >= 1 {
			A[i][j-1], A[i-1][j] = A[i-1][j], A[i][j-1]
		}
		if i+1 < 4 && j+1 < 4 {
			A[i+1][j], A[i][j+1] = A[i][j+1], A[i+1][j]
		}
	}

	return A
}