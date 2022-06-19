package jz_II

/*
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
请实现 KthLargest类：
KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。

注意：这里是第k大，而不是第k小。
*/

type KthLargest struct {
	heap minHeap //最小堆
	k    int
}

type minHeap struct {
	arr []int
}

func (a *minHeap) add(val int) {
	a.arr = append(a.arr, val)
	n := len(a.arr)
	for i := n/2 - 1; i >= 0; i-- {
		a.adjust(i, n)
	}
}

func (a *minHeap) pop() {
	a.arr = a.arr[1:]
	a.adjust(0, len(a.arr))
}

func (a *minHeap) top() int {
	return a.arr[0]
}

func (a *minHeap) size() int {
	return len(a.arr)
}

func (a *minHeap) adjust(parent, end int) {
	child := 2*parent + 1
	for child < end {
		if child+1 < end && a.arr[child+1] < a.arr[child] {
			child++
		}
		if a.arr[parent] < a.arr[child] {
			break
		}
		a.arr[parent], a.arr[child] = a.arr[child], a.arr[parent]
		parent = child
		child = 2*parent + 1
	}
}

func Constructor59(k int, nums []int) KthLargest {
	res := KthLargest{k: k}
	for _, v := range nums {
		res.Add(v)
	}
	return res
}

func (k *KthLargest) Add(val int) int {
	k.heap.add(val)
	if k.heap.size() > k.k {
		k.heap.pop()
	}
	return k.heap.top()
}
