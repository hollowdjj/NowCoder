package jz_II

/*
给定一个非负整数x，计算并返回x的平方根，即实现int sqrt(int x)函数。
正数的平方根有两个，只输出其中的正数平方根。
如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。
*/

func mySqrt(x int) int {
	//二分法
	left, right := 0, x
	res := 0
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x {
			res = mid //向下圆整，所以这里是res=mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func mySqrt1(x int) int {
	//牛顿迭代法。y1 = y0 - f(y0)/f'(y0)。 y^2 = x ---> f(y) = y^2 - x = 0
	//f'(y)=2y。f(y)/f'(y) = (y^2-x)/2y = y/2 - x/2y
	//则迭代方程为：y1 = y0/2 + x/2y0
	if x == 0 {
		return 0
	}
	res := x
	for res > x/res {
		res = (res + x/res) / 2
	}
	return res
}
