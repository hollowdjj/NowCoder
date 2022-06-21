package jz_II

import "sort"

/*
给定一个可能有重复数字的正整数数组candidates和一个目标数target，找出candidates中所有可以
使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
*/

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	res := make([][]int, 0)
	used := make([]bool, n)
	var backtrack func(curr []int, sum int, i int)
	backtrack = func(curr []int, sum int, i int) {
		if sum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}

		for j := i; j < len(candidates); j++ {
			//同层减枝
			if used[j] || j > 0 && candidates[j] == candidates[j-1] && !used[j-1] {
				continue
			}
			used[j] = true
			curr = append(curr, candidates[j])
			backtrack(curr, sum+candidates[j], j+1)
			used[j] = false
			curr = curr[:len(curr)-1]
		}
	}
	sort.Ints(candidates)
	backtrack([]int{}, 0, 0)
	return res
}
