package array

import "sort"

/*
没有重复项的数字全排列

给出一组数字，返回该组数字的所有排列
例如：
[1,2,3]的所有排列如下
[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2], [3,2,1].
（以数字在数组中的位置靠前为优先级，按字典序排列输出。）
*/

func Permute(num []int) [][]int {
	n := len(num)
	if n < 1 {
		return [][]int{num}
	}

	sort.Ints(num) //先对数字排序，从而res中的结果就是按照的字典序排列的
	res := make([][]int, 0)
	used := make([]bool, n) //used[i]表示num[i]是否在path中出现过
	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == n {
			temp := make([]int, n)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				path = append(path, num[i])
				backtrack(path)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(make([]int, 0, n))
	return res
}
