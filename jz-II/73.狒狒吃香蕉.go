package jz_II

/*
狒狒喜欢吃香蕉。这里有n堆香蕉，第i堆中有piles[i]根香蕉。警卫已经离开了，将在h小时后回来。
狒狒可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉k根。
如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉，下一个小
时才会开始吃另一堆的香蕉。狒狒喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。返回她可以在
h小时内吃掉所有香蕉的最小速度 k（k 为整数）。

输入：piles = [3,6,7,11], h = 8
输出：4
*/

func minEatingSpeed(piles []int, h int) int {
	//吃香蕉的速度和是否能在规定时间内吃完香蕉呈现出单调性，因此可以使用二分法求解
	//每个小时都必须出香蕉，因此下界为1。统计piles中的最大值max，那么上界就为max。
	//二分法找第一个
	maxVal := 0
	for _, num := range piles {
		if num > maxVal {
			maxVal = num
		}
	}
	//二分法找到最后一个不能吃完所有香蕉的速度
	left, right := 1, maxVal
	for left <= right {
		mid := (left + right) / 2
		if timecost(mid, piles) <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func timecost(speed int, piles []int) int {
	//计算以速度speed，吃完所有香蕉所需耗费的时间
	res := 0
	for _, num := range piles {
		//计算num/speed并向上取整
		res += (num + speed - 1) / speed
	}
	return res
}
