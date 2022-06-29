package jz_II

import "math/rand"

/*
设计一个支持在平均时间复杂度O(1)下，执行以下操作的数据结构：

insert(val)：当元素 val 不存在时返回 true，并向集合中插入该项，否则返回 false 。
remove(val)：当元素 val 存在时返回 true，并从集合中移除该项，否则返回 false。
getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
*/

type RandomizedSet struct {
	dic  map[int]int //数字以及其在数组中的索引
	nums []int
}

func Constructor30() RandomizedSet {
	res := RandomizedSet{dic: make(map[int]int)}
	return res
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.dic[val]; ok {
		return false
	}
	r.nums = append(r.nums, val)
	r.dic[val] = len(r.nums) - 1
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if index, ok := r.dic[val]; ok {
		last := len(r.nums) - 1
		lastVal := r.nums[last]
		r.nums[index] = lastVal
		r.dic[lastVal] = index
		r.nums = r.nums[:len(r.nums)-1]
		//必须在最后，因为val可能等于数组中的最后一个数字
		delete(r.dic, val)
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
