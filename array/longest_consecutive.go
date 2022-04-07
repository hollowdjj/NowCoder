package array

/*
最长连续序列

给定一个未排序的整数数组nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为O(n)的算法解决此问题。

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/

func LongestConsecutive(nums []int) int {
	//首先使用一个map记录所有数字
	dic := make(map[int]bool)
	for _, v := range nums {
		dic[v] = true
	}

	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		//只有当nums[i]-1不存在时，我们才去找nums[i]+1,nums[i]+2,nums[i]+3.....
		count := 1
		if !contains(dic, nums[i]-1) {
			for contains(dic, nums[i]+count) {
				count++
			}
		}
		res = max(res, count)
	}
	return res
}

func contains(dic map[int]bool, num int) bool {
	_, ok := dic[num]
	return ok
}
