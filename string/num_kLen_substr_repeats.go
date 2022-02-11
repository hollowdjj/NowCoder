package string

/*
长度为K的重复字符子串

给你一个由小写字母组成的长度为n的字符串 S ，找出所有长度为 k 且包含重复字符的子串，
请你返回全部满足要求的子串的数目。

数据范围： 2≤k≤400 ,5≤n≤900
进阶： 时间复杂度O(n)，空间复杂度O(n)
*/

func NumKLenSubstrRepeats(s string, k int) int {
	//最简单的思路是使用一个长度为k的滑动窗口，每滑动一次就判断一下
	//滑动窗口中的子串是否存在重复元素。一个更为高效的方法是，统计滑动
	//窗口中重复元素的个数。窗口滑动时，判断删除和添加了一个元素后，窗口
	//中的子串的重复元素个数。
	count := 0
	dic := make(map[byte]int)
	for i := 0; i < k; i++ {
		dic[s[i]]++
		if dic[s[i]] > 1 {
			count++
		}
	}
	res := 0
	if count > 0 {
		res++
	}
	n := len(s)
	for i := k; i < n; i++ {
		//i为滑动窗口的右边界，i-k为左边界
		dic[s[i-k]]--
		if dic[s[i-k]] == 1 {
			count--
		}
		dic[s[i]]++
		if dic[s[i]] == 2 {
			count++
		}
		if count > 0 {
			res++
		}
	}
	return res
}
