package jz_II

/*
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
*/
func findMaxLength(nums []int) int {
	//将nums中的所有0变成-1，从而问题就变成了和为0的最长连续子数组的长度
	for i, v := range nums {
		if v == 0 {
			nums[i] = -1
		}
	}
	dic := make(map[int]int) //前缀和x以及其第一次出现时的右边界
	//这里必须是-1。例如[0,1]，因为sum(i)-sum(j)表示的是nums[j+1...i]的和
	dic[0] = -1
	sum := 0
	res := 0
	for i, v := range nums {
		sum += v
		if j, ok := dic[sum]; ok {
			if t := i - j; t > res {
				res = t
			}
		} else {
			dic[sum] = i
		}
	}
	return res
}
