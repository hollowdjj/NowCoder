package jz_II

/*
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。

实现 MovingAverage 类：

MovingAverage(int size) 用窗口大小 size 初始化对象。
double next(int val)成员函数 next每次调用的时候都会往滑动窗口增加一个整数，
请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。

输入：
inputs = ["MovingAverage", "next", "next", "next", "next"]
inputs = [[3], [1], [10], [3], [5]]
输出：
[null, 1.0, 5.5, 4.66667, 6.0]

解释：
MovingAverage movingAverage = new MovingAverage(3);
movingAverage.next(1); // 返回 1.0 = 1 / 1
movingAverage.next(10); // 返回 5.5 = (1 + 10) / 2
movingAverage.next(3); // 返回 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5); // 返回 6.0 = (10 + 3 + 5) / 3
*/

type MovingAverage struct {
	slot []int
	sum  int //滑动窗口中元素和
	size int //滑动窗口大小
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	m := MovingAverage{size: size}
	return m
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.slot) < this.size {
		this.sum += val
		this.slot = append(this.slot, val)
		return float64(this.sum) / float64(len(this.slot))
	}
	this.sum -= this.slot[0]
	this.sum += val
	this.slot = this.slot[1:]
	this.slot = append(this.slot, val)
	return float64(this.sum) / float64(this.size)
}
