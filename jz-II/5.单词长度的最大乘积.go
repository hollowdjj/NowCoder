package jz_II

/*
给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的
乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回0。

输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
*/

func maxProduct(words []string) int {
	//最朴素的方法是单词两两对比，判断有无相同字符，然后在更新最大乘积。
	//使用哈希表的方式判断两个单词有无相同字符时，时间复杂度和空间复杂度
	//均为O(n)。因而，整个算法的时间复杂度为O(n^2)。考虑使用位运算来优化
	//判断两单词是否存在相同字符。一共有26个小写字母，使用一个26位的二进制
	//数，第一位表示是否有a，依次递推。判断两单词有无相同字符，只需让两个对应
	//的数字做与运算即可。
	n := len(words)
	masks := make([]int, n)
	for i, w := range words {
		for _, c := range w {
			masks[i] |= 1 << (c - 'a')
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i]&masks[j] == 0 {
				if t := len(words[i]) * len(words[j]); t > res {
					res = t
				}
			}
		}
	}
	return res
}
