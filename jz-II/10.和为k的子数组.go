package jz_II

/*
给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。

输入:nums = [1,1,1], k = 2
输出: 2
解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
*/

func subarraySum(nums []int, k int) int {
	//由于nums中的元素不全是正整数，因此不能用滑动窗口求解，而是要使用前缀和求解。
	//使用sum(i)表示nums[0..i]中所有元素之和。再使用哈希表dic记录i之前的所有前缀和。
	//若nums[0..j]的元素之和sum(j)=sum(i)-k，那么k = sum(i)-sum(j)也就是子数组
	//nums[j+1...j-1]的和
	dic := make(map[int]int)
	//前缀和为0出现的次数为1。如果dic[0]不为1，对于用例[1,2],3，当sum=3时，dic[0]=0
	dic[0] = 1
	sum, res := 0, 0
	for _, v := range nums {
		sum += v
		if count, ok := dic[sum-k]; ok {
			res += count
		}
		dic[sum]++
	}
	return res
}
