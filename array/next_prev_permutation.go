package array

import (
	"sort"
)

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
	//1. 先找到需要改变的高位：从右向左扫描排列，若一直满足nums[i] > nums[i - 1]，则说明这些元素是满足高位大于
	//低位的，不需操作，直到找到nums[i] < nums[i - 1]，找到高位比低位小的了，而且是“最低”的高位，这个位置就是我
	//们需要做交换操作的。比如：6, 8, 7, 4, 3, 2 当中的8（发现6 < 8）

	//2. 第二步找要和这个高位交换的低位：原则是尽量寻找只比这个高位“大一点”的低位，因为只是下一个排列。因为在这个
	//高位右边的数组满足从右往左递增，所以，我们重新从右边起扫描，找到第一个比这个高位大的元素。比如：6, 8, 7, 4, 3, 2 中的7.

	//3. 交换高位与低位，使得高位变大：比如6, 8, 7, 4, 3, 2 -> 7, 8, 6, 4, 3, 2

	//4. 此时，不论交换后的高位后面的元素是如何排列的，都肯定比之前的排列靠后了。因为只是下一个排列，所以我们现在
	//尽量要现在的这个排列“靠前”，怎么做呢，就是按升序排列高位后面的元素，比如：按升序排列此时7后面的元素。因为此
	//时高位后面的元素一定是从左往右，从大到小，所以，也相当于是翻转这一部分的数组。比如：7, 8, 6, 4, 3, 2 -> 7, 2, 3, 4, 6, 8
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			//从后往前找到了第一个升序对[i-1,i]。此时，找到i-1之后大于nums[i-1]的最小数nums[index]
			index := i
			for j := n - 1; j > i; j-- {
				if nums[j] > nums[i-1] {
					index = j
					break
				}
			}
			//交换nums[index]和nums[i-1]
			nums[i-1], nums[index] = nums[index], nums[i-1]
			//将i-1后面的数按照升序排序。这样才是同一前缀中的最小排列
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

func PrevPermutation(nums []int) []int {
	//[6,3,6,2,4,5,7,8]
	//1. 首先从后往前找到第一个降序对(i-1,i)，即[3,6]
	//2. 由于nums[i:]是升序排列的，所以我们从右往左找到第一个比nums[i-1]小的数nums[index]。并交换nums[i-1]和nums[index]
	//   即[6,3,5,2,4,6,7,8]。就是找比nums[i-1]小的最大值。
	//3. 最后逆序排序nums[i:]得到[6,3,5,8,7,6,4,2]
	n := len(nums)
	for i := n - 1; i > 0; i++ {
		if nums[i-1] > nums[i] {
			index := i
			for j := n - 1; j > i; j++ {
				if nums[j] < nums[j-1] {
					index = j
					break
				}
			}
			nums[i-1], nums[index] = nums[index], nums[i-1]
			//逆序排列nums[i:]
			reverse(nums, i, n-1)
			return nums
		}
	}

	reverse(nums, 0, n-1)
	return nums
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
