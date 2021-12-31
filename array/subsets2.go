package array

import "sort"

/*
给定一个整数数组 nums ，其中可能包含重复元素，请你返回这个数组的所有可能子集。
返回的答案中不能包含重复的子集，将答案按字典序进行排序。

输入：[1,2]；输出：[[],[1],[1,2],[2]]
*/

func Subsets2(nums []int) [][]int {
	//首先对数组进行排序
	sort.Ints(nums)

	//使用深度优先搜索
	var res [][]int
	var DFS func(start int, temp []int)
	DFS = func(start int, temp []int) {
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i-1] == nums[i] {
				continue
			}
			temp = append(temp, nums[i])
			DFS(i+1, temp)
			//回溯
			var t []int
			t = append(t, temp[:len(temp)-1]...)
			temp = t
		}
	}
	DFS(0, []int{})
	return res
}
