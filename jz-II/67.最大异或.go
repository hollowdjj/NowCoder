package jz_II

/*
给定一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n

输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
*/

type Trie67 struct {
	nexts [2]*Trie67 //二进制为0放nexts[0],为1放nexts[1]
}

func (t *Trie67) add(num int) {
	//注意要倒序插入，只需移动31次
	for i := 30; i >= 0; i-- {
		index := (num >> i) & 1
		if t.nexts[index] == nil {
			t.nexts[index] = &Trie67{}
		}
		t = t.nexts[index]
	}
}

func (t *Trie67) check(num int) int {
	res := 0
	for i := 30; i >= 0; i-- {
		//最大异或，就是要让二进制不同的位最多，因此要尽量选择不同的
		//位走
		index := (num >> i) & 1
		if index == 0 {
			if t.nexts[1] == nil {
				t = t.nexts[0]
				res <<= 1
			} else {
				t = t.nexts[1]
				//异或位为1
				res = res<<1 + 1
			}
		} else {
			if t.nexts[0] == nil {
				t = t.nexts[1]
				res <<= 1
			} else {
				t = t.nexts[0]
				res = res<<1 + 1
			}
		}
	}
	return res
}

func findMaximumXOR(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	res := 0
	root := &Trie67{}
	for i := 0; i < n-1; i++ {
		root.add(nums[i])
		res = max(res, root.check(nums[i+1]))
	}
	return res
}
