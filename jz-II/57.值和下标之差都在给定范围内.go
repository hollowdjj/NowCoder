package jz_II

/*
给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使
得 abs(nums[i] - nums[j]) <= t，同时又满足 abs(i - j) <= k 。

如果存在则返回 true，不存在返回 false。

输入：nums = [1,2,3,1], k = 3, t = 0
输出：true
*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	//使用桶排序，桶的大小为t。因此有：
	//[-2t-2,-t-2][-t-1,-1] [0,t] [t+1,2t] [2t+1,3t]
	//那么给定元素i，其所在的桶的ID为：
	//1. i >= 0 id = i/(t+1)
	//2. i < 0 id = (i+1)/(t+1)-1
	//若两个元素在同一个桶内，那么
	dic := make(map[int]int)
	for i, num := range nums {
		id := getID(num, t)
		if _, hit := dic[id]; hit {
			//落在同一个桶内
			return true
		}
		if prev, hit := dic[id-1]; hit && num-prev <= t {
			return true
		}
		if next, hit := dic[id+1]; hit && next-num <= t {
			return true
		}
		//向哈希表添加元素必须出现在删除的前面
		//例如[1,2],0,1。
		dic[id] = num
		if i >= k {
			delete(dic, getID(nums[i-k], t))
		}
	}
	return false
}

func getID(i int, t int) int {
	if i >= 0 {
		return i / (t + 1)
	}
	return (i+1)/(t+1) - 1
}
