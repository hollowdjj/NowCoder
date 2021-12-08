package array

/*
最长无重复子数组
给定一个长度为n的数组arr，返回arr的最长无重复连续子数组的长度
要求：空间复杂度O(n)，时间复杂度O(nlogn)
*/

//MaxLength 滑动窗口解决最长无重复连续子数组问题
func MaxLength(arr []int) int {
	m := make(map[int]int)
	res := 0
	left, right := -1, 0
	for ; right < len(arr); right++ {
		if i, ok := m[arr[right]]; ok {
			//当出现重复元素时，我们需要比较一下left和该元素首次出现的索引i
			//的大小，然后取最大的那个，以缩小窗口。
			left = MaxOfTwoInt(left, i)
		}
		//更新窗口的大小
		res = MaxOfTwoInt(res, right-left)
		//更新重复数字的最新索引
		m[arr[right]] = right
	}
	return res
}

func MaxOfTwoInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinOfTwoInt(a, b int) int {
	if a < b {
		return a
	}
	return b

}
