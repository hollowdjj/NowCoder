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
			//只能按一个顺序进行访问连续相同的元素，才能保证唯一。以[1,1,1,2]为例，我们想要得到正确的解
			//正确的访问顺序为：
			//1. 从num[0]开始，一直往后走得到[1,1,1,2]
			//2. 从num[1]开始，一直往后走得到[1,1,2,1]
			//3. 从num[2]开始，一直往后走得到[1,2,1,1]
			//4. 从num[3]开始，一直往后走得到[2,1,1,1]
			//但是我们不能简单的用下面这个循环完成，因为题目要求结果必须以字典序排列。下面这段代码在输入
			//为[2,2,-1]时，输出为[[-1,2,2],[2,2,-1],[2,-1,2]]，而不是[[-1,2,2],[2,-1,2],[2,2,-1]]
			//for i:=0;i<n;i++ {
			//	path := make([]int,0,n)
			//	path = append(path,num[i:]...)
			//	path = append(path,num[:i]...)
			//	res = append(res,path)
			//}
			//所以下面这个条件判断就是从
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
