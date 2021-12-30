package array

import (
	"sort"
)

/*
加起来和为目标值的组合(二)
给出一组候选数c和一个目标数t，找出候选数中加起来和等于t的所有组合。c中的每个数字在一个组合中只能使用一次。

注意：
1. 题中的所有数字，包括目标数t都是正整数
2. 组合中的数字要按递增顺序排列
3. 结果中不能包含重复的组合
4. 组合之间的排序按照索引从小到大依次比较，小的排在前面，如果索引相同的情况下数字相同，则比较下一个索引
数据范围：1≤n≤70

示例：
输入：[100,10,20,70,60,10,50],80; 返回：[[10,10,60],[10,20,50],[10,70],[20,60]]
*/

var combinations [][]int

func CombinationSum2(num []int, target int) [][]int {
	sort.Ints(num)
	var temp []int
	backTrack(num, 0, temp, target)
	return combinations
}

func backTrack(num []int, start int, temp []int, target int) {
	if target == 0 {
		combinations = append(combinations, temp)
		return
	}

	if start >= len(num) {
		return
	}

	for i := start; i < len(num); i++ {
		//去重(与和为target的三元组一样)
		if i > start && num[i] == num[i-1] {
			continue
		}
		//剪枝
		if num[i] > target {
			continue
		}
		temp = append(temp, num[i])
		backTrack(num, i+1, temp, target-num[i])
		t := make([]int, len(temp)-1)
		copy(t, temp[:len(temp)-1])
		temp = t
	}
}
