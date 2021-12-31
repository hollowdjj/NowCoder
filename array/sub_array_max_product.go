package array

import "math"

/*
给定一个double类型的数组arr，其中的元素可正可负可0，返回连续子数组累乘的最大乘积。

数据范围：数组大小满足0≤n≤10，数组中元素满足∣val∣≤10
进阶：空间复杂度 O(1)，时间复杂度O(n)

输入：[-2.5,4,0,3,0.5,8,-1]
输出：12.00000(取连续子数组[3,0.5,8]可得累乘的最大乘积为12.00000)
*/

func SubArrayMaxProduct(arr []float64) float64 {
	//此题由于负数和负数相乘为正数，故需要两个dp数组。一个记录以i-1结尾的
	//最大连续乘积，另一个记录最小连续乘积。故状态转移方程为：
	//    dpMax[i] = Max{arr[i],arr[i]*dpMax[i],arr[i]*dpMin[i]}
	//    dpMin[i] = Min{arr[i],arr[i]*dpMax[i],arr[i]*dpMin[i]}
	min, max, res := arr[0], arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		lastMin, lastMax := min, max
		max = math.Max(arr[i], math.Max(arr[i]*lastMax, arr[i]*lastMin))
		min = math.Min(arr[i], math.Min(arr[i]*lastMax, arr[i]*lastMin))
		res = math.Max(max, res)
	}
	return res
}
