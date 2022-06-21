package jz_II

/*
给你一个整数数组nums，除某个元素仅出现一次外，其余每个元素都恰出现三次。
请你找出并返回那个只出现了一次的元素。

1 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
*/

func singleNum(nums []int) int {
	count := make([]int32, 32)
	for _, v := range nums {
		t := int32(v)
		for i := 0; i < 32; i++ {
			count[i] += (t >> i) & 1
		}
	}

	var res int32
	for i := 31; i >= 0; i-- {
		res |= count[i] % 3 << i
	}
	return int(res)
}
