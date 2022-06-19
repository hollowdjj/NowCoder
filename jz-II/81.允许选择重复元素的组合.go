package jz_II

/*
给定一个无重复元素的正整数数组candidates和一个正整数target，找出candidates中所有可以使数字
和为目标数target的唯一组合。
candidates中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是不同的。
对于给定的输入，保证和为target 的唯一组合数少于 150 个。

输入: candidates = [2,3,6,7], target = 7
输出: [[7],[2,2,3]]
*/

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	res := make([][]int, 0)
	var backtrack func(i int, curr []int, sum int)
	backtrack = func(i int, curr []int, sum int) {
		if sum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}
		for j := i; j < n; j++ {
			curr = append(curr, candidates[j])
			backtrack(j, curr, sum+candidates[j])
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0, []int{}, 0)
	return res
}
