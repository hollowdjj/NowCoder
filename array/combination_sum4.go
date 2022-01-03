package array

/*
给定一个由不同整数构成的数组 nums 和一个整数 target ，请你从 nums 找出总和是 target 的组合的个数。
解集中可以重复使用 nums 中的元素。且解集中数字顺序不同视为不同的组合。

输入：[1,2,3],4
返回值：7
说明：(1,1,1,1) (1,1,2) (1,2,1) (1,3) (2,1,1) (2,2) (3,1)
*/

func CombinationSum4(nums []int, target int) int {
	totalCount := 0
	var bt func(rest int)
	bt = func(rest int) {
		if rest < 0 {
			return
		} else if rest == 0 {
			totalCount++
		}

		for i := 0; i < len(nums); i++ {
			bt(rest - nums[i])
		}
	}
	bt(target)
	return totalCount
}
