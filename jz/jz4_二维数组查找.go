package jz

/*
在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
给定 target = 7，返回 true。
给定 target = 3，返回 false。

数据范围：矩阵的长宽满足0≤n,m≤500，矩阵中的值满足0≤val≤10^9
进阶：空间复杂度O(1) ，时间复杂度 O(n+m)
*/

func Find(target int, array [][]int) bool {
	n, m := len(array), len(array[0])
	i, j := 0, m-1
	//从右上角开始搜索
	for i < n && j >= 0 {
		if array[i][j] > target {
			j--
		} else if array[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}
