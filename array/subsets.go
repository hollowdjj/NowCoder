package array

import "sort"

/*
现在有一个没有重复元素的整数集合S，求S的所有子集
注意：
你给出的子集中的元素必须按升序排列
给出的解集中不能出现重复的元素

数据范围：1≤n≤5，集合中的任意元素满足∣val∣≤10
要求：空间复杂度 O(n!)，时间复杂度 O(n!)

输入：[1,2,3] 输出：[[],[1],[2],[3],[1,2],[1,3],[2,3],[1,2,3]]


void DFS(参数) {
	if(满足存放结果的条件) {
		存放结果
		return
	}

	for(选择 : 本层集合中的元素) {
		处理当前节点
		DFS(下一节点)
		回溯
	}
}
*/

func SubSets(A []int) [][]int {
	//先使用深度优先搜索
	sort.Ints(A)
	var res [][]int
	var dfs func(level int, temp []int)
	dfs = func(level int, temp []int) {
		res = append(res, temp)
		for i := level; i < len(A); i++ {
			temp = append(temp, A[i]) //选第i个
			dfs(i+1, temp)
			t := make([]int, len(temp)-1) //不选第i个
			copy(t, temp[:len(temp)-1])
			temp = t
		}
	}
	dfs(0, []int{})

	//由于题目要求的是按照解集中元素的个数排列，相同个数的则通过首元素大小排序故最后还要进行一下排序
	for i := 0; i < len(res)-1; i++ {
		if len(res[i+1]) > len(res[i]) {
			//从i+1往后找到第一个长度等于len(res[i])的那个数组的索引index
			index := 0
			for j := i + 2; j < len(res); j++ {
				if len(res[j]) == len(res[i]) {
					index = j
					break
				}
			}
			//将第index个数组交换到第i个数组后面
			for j := index; j > i+1; j-- {
				temp := res[j-1]
				res[j-1] = res[j]
				res[j] = temp
			}
		}
	}
	return res
}
