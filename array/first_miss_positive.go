package array

/*
缺失的第一个正数

给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/

func FirstMissingPositive(nums []int) int {
	//原地哈希，将元素大小为nums[i]的元素，放置在nums[i]-1的位置上
	//那么我们如何将数组进行恢复呢？我们可以对数组进行一次遍历，对于遍历到的数x=nums[i]
	//如果x∈[1,N]，我们就知道x应当出现在数组中的x−1的位置，因此交换nums[i]和nums[x−1]，
	//这样x就出现在了正确的位置。在完成交换后，新的nums[i]可能还在[1,N]的范围内，我们需要继续
	//进行交换操作，直到x不属于[1,N]。
	//从而当nums[i] != i+1时，i+1就是缺失的第一个正数
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
