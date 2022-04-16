package jz

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*/
func printNumbers(n int) []int {
	limit := 9
	for n > 1 {
		limit = limit*10 + 9
		n--
	}

	res := make([]int, limit)
	for i := 1; i <= limit; i++ {
		res[i-1] = i
	}
	return res
}
