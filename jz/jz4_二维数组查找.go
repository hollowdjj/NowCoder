package jz

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
