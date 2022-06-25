package jz_II

/*
给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
*/
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
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
			if used[i] {
				continue
			}
			used[i] = true
			curr = append(curr, nums[i])
			backtrack(curr)
			used[i] = false
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{})
	return res
}
