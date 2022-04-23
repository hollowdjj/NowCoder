package jz

/*
求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/

func Sum(n int) int {
	res := 0
	var work func(n int) bool
	work = func(n int) bool {
		res += n
		return n > 0 && work(n-1)
	}
	work(n)
	return res
}
