package array

/*
矩阵元素查找

已知一个有序矩阵mat，同时给定矩阵的大小n和m以及需要查找的元素x，且矩阵的行和列都是从小到大有序的。
设计查找算法返回所查找元素的二元数组，代表该元素的行号和列号(均从零开始)。保证元素互异。

要求：空间复杂度 O(1)O(1)，时间复杂度 O(n+m)O(n+m)
*/

func FindElement(mat [][]int, n int, m int, x int) []int {
	i := 0
	bound := m
	res := make([]int, 2)
	for ; i < n; i++ {
		if mat[i][0] > x || mat[i][m-1] < x {
			continue
		}
		if mat[i][m-1] >= x {
			//使用二分法找到mat[i][0:bound]中第一个大于等于x的元素的索引
			j := find(mat[i][:bound], x)
			if mat[i][j] == x {
				res[0] = i
				res[1] = j
				break
			} else {
				bound = j
			}
		}
	}

	return res
}

//在arr中找到第一个大于等于x的元素索引。找到
func find(arr []int, x int) int {
	n := len(arr)
	left, right := 0, n-1
	for left < right {
		mid := (right + left) / 2
		if arr[mid] < x {
			left = mid + 1
		} else if arr[mid] > x {
			right = mid - 1
		} else {
			return mid
		}
	}

	if arr[left] < x {
		left++
	}
	return left
}
