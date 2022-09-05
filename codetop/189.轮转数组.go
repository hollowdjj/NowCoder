package codetop

/*
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
*/

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse := func(nums []int) {
		i, j := 0, len(nums)-1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j-1
		}
	}
	//变为7 6 5 4 3 2 1
	reverse(nums)
	//翻转前k个得5 6 7 4 3 2 1
	reverse(nums[:k])
	//翻转后4个得5 6 7 1 2 3 4
	reverse(nums[k:])
}
