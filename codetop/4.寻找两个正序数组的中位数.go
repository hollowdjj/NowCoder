package codetop

import "math"

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//当n1+n2为奇数时，(n1+n2+1)/2,(n1+n2+2)/2的值相等
	//当n1+n2为偶数时，(n1+n2+1)/2 = (n1+n2+2)/2 - 1
	n1, n2 := len(nums1), len(nums2)
	left, right := (n1+n2+1)/2, (n1+n2+2)/2
	res := findKth(nums1, nums2, 0, 0, left) + findKth(nums1, nums2, 0, 0, right)
	return float64(res) / 2
}

//在两个有序数组nums1[i:]、nums2[j:]表示的有序数组中找到第k个元素
func findKth(nums1, nums2 []int, i, j, k int) int {
	if i >= len(nums1) {
		//直接在第二个数组中找
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		//直接在第一个数组中找
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	//这里取最大值是因为，若某数组的长度小于k/2，那么另一个数组的前k/2个元素中也
	//一定不存在中位数。因为，最好的情况也是第一个数组中的所有元素均小于另一个数组
	//的首元素，即使这样长度也是小于k的。
	m1, m2 := math.MaxInt32, math.MaxInt32
	if i+k/2-1 < len(nums1) {
		m1 = nums1[i+k/2-1]
	}
	if j+k/2-1 < len(nums2) {
		m2 = nums2[j+k/2-1]
	}
	//nums1的第k/2个数较小，那么说明要找的数字肯定不在num1的前k/2个数字中
	if m1 < m2 {
		return findKth(nums1, nums2, i+k/2, j, k-k/2)
	}
	return findKth(nums1, nums2, i, j+k/2, k-k/2)
}
