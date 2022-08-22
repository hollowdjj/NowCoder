package codetop

import "sort"

/*
给你两个数组，arr1 和 arr2，arr2 中的元素各不相同，arr2 中的每个元素都出现在 arr1 中。
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2
中出现过的元素需要按照升序放在 arr1 的末尾。

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]
*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dic := make(map[int]int)
	for i, num := range arr2 {
		dic[num] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		ri, hasi := dic[arr1[i]]
		rj, hasj := dic[arr1[j]]
		if hasi && hasj {
			//两个数在arr2中都存在，那么就按照在arr2中的索引排序，保持相对顺序
			return ri < rj
		}
		if hasi || hasj {
			//一个出现在arr2中，另一个没有出现在arr2中。出现在arr2中的在前
			return hasi
		}
		//两个数在arr2中均不存在，那么按照升序排列在最后
		return arr1[i] < arr1[j]
	})
	return arr1
}
