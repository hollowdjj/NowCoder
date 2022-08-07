package jz_II

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/

func longestConsecutive(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	res := 0
	//当dic[num-1]存在时，无需找以num开头连续序列，因为肯定是被遍历过了的
	for _, num := range nums {
		if _, hit := dic[num-1]; !hit {
			count := 0
			for {
				num++
				count++
				if _, ok := dic[num]; !ok {
					break
				}
			}
			res = max(res, count)
		}
	}
	return res
}
