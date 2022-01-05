package array

import "sort"

/*
给定一个数组，将数组重新排列，得到一系列数组排列S，请你从S中，找出恰好比当前数组排列字典序大于1的数组排列。
1.[1,2,3]的得到的数组排列S有:[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]。
2.该题数组排列的字典序大小排序规则：2个数组排列的元素按顺序比较，直到数组元素不相等为止，不相等的第一个元素，
  谁的元素大，谁的字典序比较大，比如数组a=[1,2,3]与数组b=[1,3,2]比较：a[0]=b[0]，a[1]<b[1]，此时出现了第
  一个不相同的，且a[1]<b[1]，则a的字典序小于b的字典序。且[1,3,2]的字典序在排列S中，正好在[1,2,3]的后面，视
  为[1,3,2]的字典序比[1,2,3]的字典序大于1。
3.如果不存在更大的数组排列，则返回最小的数组排列。
*/

func NextPermutation(nums []int) []int {
	//具体的思路为：首先，从后往前，找到第一个升序对(i,i+1)。然后，在i后面找到大于nums[i]的最小数nums[j]，
	//并交换nums[i]和nums[j]。最后对i后面的数安装升序排列即可。
	//这里解释一下，为什么要从后往前找第一个升序对。以[1,4,2,3]为例，比它刚好大1的排序为[1,4,3,2]。如果，
	//是从前往后找，那么结果就是[1,2,3,4]。
	n := len(nums)
	for i := n - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			//从后往前找到了第一个升序对[i-1,i]。此时，找到i-1之后大于nums[i]的最小数nums[index]
			min := nums[i]
			index := i
			for j := i; j < n; j++ {
				if nums[j] < min {
					min = nums[j]
					index = j
				}
			}
			//交换nums[index]和nums[i-1]
			nums[i-1], nums[index] = nums[index], nums[i-1]
			//将i-1后面的数安装升序排序
			sort.Ints(nums[i:])
			return nums
		}
	}
	//如果一个升序对都没找到，说明已经是最大排列，此时需要返回最小排列，即当前数组的逆序
	left, right := 0, n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
