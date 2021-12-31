package array

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

数据范围：矩阵的长宽满足 0≤n,m≤500 ， 矩阵中的值满足 0≤val≤10^9
进阶：空间复杂度 O(1) ，时间复杂度 O(n+m)
*/

func Find(target int, array [][]int) bool {
	n, m := len(array), len(array[0])
	//首先遍历第一行，找到第一个大于或等于target的数的索引index
	index := 0
	for ; index < m; index++ {
		if array[0][index] == target {
			return true
		}
		if array[0][index] > target {
			break
		}
	}
	if index == 0 {
		return false
	}

	for i := 1; i < n; i++ {
		//从上之下遍历第index-1列
		if array[i][index-1] == target {
			return true
		}
		//如果第index-1列没找到target，需要跳到左边一列的第一个元素
		if array[i][index-1] > target {
			if (index - 1) == 0 {
				break
			}
			index--
			i-- //这里必须i--保证是从左边一列的同一行往下搜索
		}
	}

	return false
}
