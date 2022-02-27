package array

import "sort"

/*
有重复项数字的全排列

给出一组可能包含重复项的数字，返回该组数字的所有排列。结果以字典序升序排列。

输入：[1,1,2]
返回：[[1,1,2],[1,2,1],[2,1,1]]
*/

func PermuteUnique(num []int) [][]int {
	n := len(num)
	if n < 1 {
		return [][]int{num}
	}

	//先排序一下数组
	sort.Ints(num)
	used := make([]bool, n)
	res := make([][]int, 0)
	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == n {
			temp := make([]int, n)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			//在两种情况下，我们需要跳过i：
			//1. path中使用了第i个数字
			//2. path没有使用第i个数字，但是使用了第i-1个数字，且这两个数字相等。使用了第i-1个数字后，used[i-1]==false
			//所以这里是!used[i-1]
			if used[i] || i > 0 && num[i] == num[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, num[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(make([]int, 0, n))
	return res
}
