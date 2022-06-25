package jz_II

import "sort"

/*
给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
*/

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	used := make([]bool, n)
	var backtrack func(curr []int)
	backtrack = func(curr []int) {
		if len(curr) == n {
			temp := make([]int, n)
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] || i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			used[i] = true
			curr = append(curr, nums[i])
			backtrack(curr)
			used[i] = false
			curr = curr[:len(curr)-1]
		}
	}
	sort.Ints(nums)
	backtrack([]int{})
	return res
}
