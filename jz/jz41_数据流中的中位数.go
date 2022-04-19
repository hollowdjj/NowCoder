package jz

/*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取
数据流，使用GetMedian()方法获取当前读取数据的中位数。

输入：[5,2,3,4,1,6,7,0,8]
输出："5.00 3.50 3.00 3.50 3.00 3.50 4.00 3.50 4.00"
说明：数据流里面不断吐出的是5,2,3...,则得到的平均数分别为5,(5+2)/2,3...
*/

var minHeap []int //小根堆，维护数据流中的较大值
var maxHeap []int //大根堆，维护数据流中的较小值
var streamNum int

func insertToMinHeap(num int) {
	minHeap = append(minHeap, num)
	n := len(minHeap)
	for i := n/2 - 1; i >= 0; i-- {
		adjustMinHeap(i, n)
	}
}

func popMinHeap() {
	minHeap = minHeap[1:]
	adjustMinHeap(0, len(minHeap))
}

func insertToMaxHeap(num int) {
	maxHeap = append(maxHeap, num)
	n := len(maxHeap)
	for i := n/2 - 1; i >= 0; i-- {
		adjustMaxHeap(i, n)
	}
}

func popMaxHeap() {
	maxHeap = maxHeap[1:]
	adjustMaxHeap(0, len(maxHeap))
}

func adjustMinHeap(parent, end int) {
	child := parent*2 + 1
	for child < end {
		if child+1 < end && minHeap[child+1] < minHeap[child] {
			child++
		}
		if minHeap[parent] <= minHeap[child] {
			break
		}
		minHeap[parent], minHeap[child] = minHeap[child], minHeap[parent]
		parent = child
		child = 2*parent + 1
	}
}

func adjustMaxHeap(parent, end int) {
	child := 2*parent + 1
	for child < end {
		if child+1 < end && maxHeap[child+1] > maxHeap[child] {
			child++
		}
		if maxHeap[parent] >= maxHeap[child] {
			break
		}
		maxHeap[parent], maxHeap[child] = maxHeap[child], maxHeap[parent]
		parent = child
		child = 2*parent + 1
	}
}

func Insert(num int) {
	//使用一个小根堆和一个大根堆。小根堆维护数据流中的较大值，大根堆维护数据流中的较小值。
	//1. 当数据个数为奇数时，我们选择将num插入到大根堆中。由于大根堆中维护的是较小值，所以实际插入到大根堆中的元素
	//   为num与小根堆堆顶两者的较小值。
	//2. 当数据个数为偶数时，我们选择将num插入到大顶堆中。由于小根堆中维护的是较大值，所以实际插入到小根堆中的元素
	//	 为num与大根堆堆顶两者的较大值。
	//从而当元素个数为奇数时，中间值就是大顶堆堆顶；为偶数时，就是大顶堆和小顶堆堆顶元素的平均值。
	streamNum++
	if streamNum%2 == 1 {
		if len(minHeap) > 0 && minHeap[0] < num {
			//将num存入小顶堆，并将小顶堆堆顶元素放到大顶堆中
			insertToMinHeap(num)
			num = minHeap[0]
			popMinHeap()
		}
		insertToMaxHeap(num)
	} else {
		if len(maxHeap) > 0 && maxHeap[0] > num {
			insertToMaxHeap(num)
			num = maxHeap[0]
			popMaxHeap()
		}
		insertToMinHeap(num)
	}
}

func GetMedian() float64 {
	if streamNum%2 == 1 {
		return float64(maxHeap[0])
	} else {
		return (float64(maxHeap[0]) + float64(minHeap[0])) / 2
	}
}
