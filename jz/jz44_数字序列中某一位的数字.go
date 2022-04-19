package jz

import "math"

/*
数字以 0123456789101112131415... 的格式作为一个字符序列，在这个序列中第 2 位（从下标 0 开始计算）是 2 ，第 10 位是 1
第 13 位是 1 ，以此类题，请你输出第 n 位对应的数字。
*/

func findNthDigit(n int) int {
	//0~9，1位数字，一共有9个
	//10~99, 2位数字，一共有90个
	//100~999, 3位数字，一共有900个
	//因此，首先对n进行递减，判断第n位所在的数字是几位数字
	d := 1
	for count := 9; d*count < n; count *= 10 {
		n -= d * count
		d++
	}

	//计算d位数字的最小值
	start := int(math.Pow10(d - 1))
	//用剩余的n除以d得到偏移量
	offset := n / d
	//得到第n个字符所在的目标数字
	target := start + offset - 1

	n -= d * offset
	if n == 0 {
		//说明n/d刚好除尽，那么第n位就是target的最后一位数字
		return target % 10
	}
	//没有除尽，那么答案就是target+1的第n位数字121
	return (target + 1) / int(math.Pow10(d-n)) % 10
}
