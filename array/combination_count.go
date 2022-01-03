package array

/*
给定一个无重复元素的正整数数组 nums 和一个正整数 target ，找出 nums 中所有可以使数字之和为目标数
target 的组合，nums 中的数可以重复选取，只要选出的组合中有一个数不同则视为是不同组合，但是数字相同
顺序不同的组合不能视为不同的组合

输入：5,[1,4,5]
输出：[[1,4],[5],[1,1,1,1,1]]
*/

func CombinationCount(target int, nums []int) [][]int {
	var res [][]int

	var dfs func(i int, rest int, temp []int)
	dfs = func(i int, rest int, temp []int) {
		if rest < 0 {
			return
		} else if rest == 0 {
			res = append(res, temp)
		}
		for j := i; j < len(nums); j++ {
			temp = append(temp, nums[j])
			//一个数字可以重复使用，所以这里是j而不是j+1
			//要是顺序不同，也算不同解，那么每次j都从0开始即可
			dfs(j, rest-nums[j], temp)
			t := make([]int, len(temp)-1)
			copy(t, temp[:len(temp)-1])
			temp = t
		}
	}
	dfs(0, target, []int{})
	return res
}
