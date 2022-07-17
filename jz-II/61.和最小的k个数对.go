package jz_II

/*
给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := make([][]int, 0)
	heap := maxHeap61{}
	n1, n2 := len(nums1), len(nums2)
	for i := 0; i < n1 && i < k; i++ {
		for j := 0; j < n2 && j < k; j++ {
			if len(heap.arr) < k {
				heap.add([]int{nums1[i], nums2[j]})
			} else {
				if heap.arr[0][0]+heap.arr[0][1] > nums1[i]+nums2[j] {
					heap.pop()
					heap.add([]int{nums1[i], nums2[j]})
				}
			}
		}
	}
	for _, s := range heap.arr {
		res = append(res, s)
	}
	return res
}

type maxHeap61 struct {
	arr  [][]int
	nmax int
}

func (m *maxHeap61) add(pair []int) {
	m.arr = append(m.arr, pair)
	n := len(m.arr)
	for i := n/2 - 1; i >= 0; i-- {
		m.adjust(i, n)
	}
}

func (m *maxHeap61) pop() []int {
	res := m.arr[0]
	m.arr = m.arr[1:]
	m.adjust(0, len(m.arr))
	return res
}

func (m *maxHeap61) adjust(parent, end int) {
	child := 2*parent + 1
	for child < end {
		if child+1 < end && m.arr[child+1][0]+m.arr[child+1][1] >
			m.arr[child][0]+m.arr[child][1] {
			child++
		}
		if m.arr[parent][0]+m.arr[parent][1] >= m.arr[child][0]+m.arr[child][1] {
			break
		}
		m.arr[parent], m.arr[child] = m.arr[child], m.arr[parent]
		parent = child
		child = 2*parent + 1
	}
}
