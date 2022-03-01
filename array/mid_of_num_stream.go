package array

/*
数据流中的中位数

如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取
数据流，使用GetMedian()方法获取当前读取数据的中位数。
*/

var minHeap []int //小顶堆，存放较大数
var maxHeap []int //大顶堆，存放较小数。从而大顶顶堆堆顶是较小数中的最大数，而小顶堆的堆顶是较大数中的最小数，从而就能得到中位数
var streamNum int

func Insert(num int) {
	streamNum++

	//如果是奇数的话，我们选择把数放入大顶堆。但由于大顶堆存放的是较小数，所以
	//我们需要判断一下num和小顶堆堆顶元素的大小关系。
	if streamNum%2 == 1 {
		if len(minHeap) > 0 && num > minHeap[0] {
			//把num存入小顶堆，然后把小顶堆的堆顶元素存入大顶堆中
			insertMinHeap(num)
			num = minHeap[0]
			popMinHeap()
		}
		insertMaxHeap(num)
	} else {
		//如果是偶数，那么选择插小顶堆。
		if len(maxHeap) > 0 && num < maxHeap[0] {
			//把num存入大顶堆，然后把大顶堆的堆顶元素存入小顶堆中
			insertMaxHeap(num)
			num = maxHeap[0]
			popMaxHeap()
		}
		insertMinHeap(num)
	}
}

func GetMedian() float64 {
	if streamNum%2 == 1 {
		//如果是奇数，那么中位数就是大顶堆的堆顶元素
		return float64(maxHeap[0])
	}
	return (float64(minHeap[0]) + float64(maxHeap[0])) / 2.0
}

func insertMinHeap(num int) {
	minHeap = append(minHeap, num)
	for i := len(minHeap)/2 - 1; i >= 0; i-- {
		adjustMinHeap(minHeap, i, len(minHeap))
	}
}

func popMinHeap() {
	minHeap = minHeap[1:]
	adjustMinHeap(minHeap, 0, len(minHeap))
}

func popMaxHeap() {
	maxHeap = maxHeap[1:]
	adjustMaxHeap(maxHeap, 0, len(maxHeap))
}

func insertMaxHeap(num int) {
	maxHeap = append(maxHeap, num)
	for i := len(maxHeap)/2 - 1; i >= 0; i-- {
		adjustMaxHeap(maxHeap, i, len(maxHeap))
	}
}

func adjustMinHeap(arr []int, parentIndex int, end int) {
	childIndex := 2*parentIndex + 1

	for childIndex < end {
		//选取左右孩子中的较小值
		if childIndex+1 < end && arr[childIndex+1] < arr[childIndex] {
			childIndex++
		}
		if arr[parentIndex] < arr[childIndex] {
			break
		}
		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
}

func adjustMaxHeap(arr []int, parentIndex int, end int) {
	childIndex := 2*parentIndex + 1
	for childIndex < end {
		if childIndex+1 < end && arr[childIndex+1] > arr[childIndex] {
			childIndex++
		}
		if arr[parentIndex] > arr[childIndex] {
			break
		}
		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
}
