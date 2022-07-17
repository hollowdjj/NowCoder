package jz_II

/*
给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
*/

func topKFrequent(nums []int, k int) []int {
	//哈希表记录元素出现次数
	dic := make(map[int]int)
	for _, n := range nums {
		dic[n]++
	}
	//维护一个大小为k的小根堆
	heap := minHeap60{nmax: k}
	for k, v := range dic {
		heap.add60([2]int{k, v})
	}
	res := make([]int, 0, k)
	for _, n := range heap.arr {
		res = append(res, n[0])
	}
	return res
}

type minHeap60 struct {
	arr  [][2]int
	nmax int
}

func (m *minHeap60) add60(val [2]int) {
	if len(m.arr) < m.nmax {
		m.arr = append(m.arr, val)
		//调整为小根堆
		n := len(m.arr)
		for i := n - 1; i >= 0; i-- {
			m.adjust60(i, n)
		}
	} else {
		if val[1] <= m.arr[0][1] {
			return
		}
		m.arr[0] = val
		m.adjust60(0, m.nmax)
	}
}

func (m *minHeap60) adjust60(parent, end int) {
	child := 2*parent + 1
	for child < end {
		//找到左右子节点中的较小值
		if child+1 < end && m.arr[child+1][1] < m.arr[child][1] {
			child++
		}
		if m.arr[parent][1] <= m.arr[child][1] {
			break
		}
		m.arr[parent], m.arr[child] = m.arr[child], m.arr[parent]
		parent = child
		child = 2*parent + 1
	}
}
