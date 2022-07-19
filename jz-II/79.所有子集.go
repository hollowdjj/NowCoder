package jz_II

/*
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	var dfs func(path []int, i int)
	dfs = func(path []int, i int) {
		if i == len(nums) {
			return
		}
		for j := i; j < len(nums); j++ {
			path = append(path, nums[j])
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			dfs(path, j+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
