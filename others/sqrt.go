package others

import "math"

/*
求平方根
*/

func Sqrt(x int) int {
	//牛顿迭代法：x_k+1 = x_k - f(x_k)/f'(x_k)
	//求x的平方根，其实就是求方程f(y)=y^2-x=0 f'(y) = 2y。因此，迭代公式为y1=y0-(y0^2-x)/2y0
	//即y1=0.5(y0+x/y0)
	var err float64 = 1
	var last float64 = 0 //上一次迭代的结果
	var res float64 = 1  //本次迭代的结果
	for math.Abs(res-last) >= err {
		last = res
		res = (res + float64(x)/res) / 2
	}
	return int(res) //向下取整就是截断
}
