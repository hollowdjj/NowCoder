package array

import "math"

/*
三个数的最大乘积
给定一个长度为n的无序数组A，包含正数、负数和0，请从中找出3个数，使得乘积最大，返回这个乘积。
要求时间复杂度：O(n)，空间复杂度O(1)
数据范围：3≤n
*/

//MaxProduct 无序数组中三个数的最大乘积
func MaxProduct(A []int) int64 {
	//思路就是遍历一次数组，找到最大的三个数，以及最小的两个数
	smallest1, smallest2 := math.MaxInt, math.MaxInt
	biggest1, biggest2, biggest3 := math.MinInt, math.MinInt, math.MinInt

	for _, v := range A {
		if v > biggest3 {
			if v > biggest2 {
				if v > biggest1 {
					biggest3 = biggest2
					biggest2 = biggest1
					biggest1 = v
				} else {
					biggest3 = biggest2
					biggest2 = v
				}
			} else {
				biggest3 = v
			}
		}

		if v < smallest2 {
			if v < smallest1 {
				smallest2 = smallest1
				smallest1 = v
			} else {
				smallest2 = v
			}
		}
	}

	val1 := smallest1 * smallest2 * biggest1
	val2 := biggest1 * biggest2 * biggest3
	res := 0
	if val1 < val2 {
		res = val2
	} else {
		res = val1
	}
	return int64(res)
}
