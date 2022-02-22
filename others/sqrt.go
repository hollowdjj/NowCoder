package others

/*
求平方根
实现函数 int sqrt(int x).
计算并返回 x 的平方根（向下取整）

要求：空间复杂度O(1)，时间复杂度O(logx)
*/

//牛顿法求向下取整的平方根
func Sqrt(x int) int {
	//牛顿迭代法：x_k+1 = x_k - f(x_k)/f'(x_k)
	//求x的平方根，其实就是求方程f(y)=y^2-x=0 f'(y) = 2y。因此，迭代公式为y1=y0-(y0^2-x)/2y0
	//即y1=0.5(y0+x/y0)
	res := x
	for res > x/res {
		res = (res + x/res) / 2
	}
	return int(res) //向下取整就是截断
}

//二分法求向下取整的平方根
func SqrtAdvance(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	s := x / 2
	for s*s > x {
		s /= 2
	}
	for s*s <= x {
		s++
	}
	return s - 1
}
