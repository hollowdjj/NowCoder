package jz_II

/*
给定一个长度为n的整数数组nums，其中nums是范围为[1，n]的整数的排列。还提供了一个2D整数数组 sequences
其中sequences[i]是nums的子序列。检查nums是否是唯一的最短超序列 。最短超序列是长度最短的序列，并且所有
序列sequences[i]都是它的子序列。对于给定的数组sequences，可能存在多个有效的超序列 。

例如，对于 sequences = [[1,2],[1,3]] ，有两个最短的 超序列 ，[1,2,3] 和 [1,3,2] 。
而对于 sequences = [[1,2],[1,3],[1,2,3]] ，唯一可能的最短 超序列 是 [1,2,3] 。[1,2,3,4] 是可能的超序列，但不是最短的。
如果 nums 是序列的唯一最短 超序列 ，则返回 true ，否则返回 false 。
子序列 是一个可以通过从另一个序列中删除一些元素或不删除任何元素，而不改变其余元素的顺序的序列。

输入：nums = [1,2,3], sequences = [[1,2],[1,3]]
输出：false
解释：有两种可能的超序列：[1,2,3]和[1,3,2]。
序列 [1,2] 是[1,2,3]和[1,3,2]的子序列。
序列 [1,3] 是[1,2,3]和[1,3,2]的子序列。
因为 nums 不是唯一最短的超序列，所以返回false
*/

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	//建图并统计节点的入度
	n := len(nums) + 1
	graph := make([]map[int]bool, n) //使用哈希表作为邻接链表避免重复
	for i := 1; i < n; i++ {
		graph[i] = make(map[int]bool)
	}
	indeg := make([]int, n)
	for i := range sequences {
		arr := sequences[i]
		for j := 0; j < len(arr)-1; j++ {
			graph[arr[j]][arr[j+1]] = true
			indeg[arr[j+1]]++
		}
	}
	//将所有入度为0的节点放入队列
	var q []int
	for i := 1; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		if len(q) != 1 {
			return false
		}
		front := q[0]
		q = q[1:]
		for node := range graph[front] {
			indeg[node]--
			if indeg[node] == 0 {
				q = append(q, node)
			}
		}
	}
	return true
}
