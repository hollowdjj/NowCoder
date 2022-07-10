package jz_II

/*
给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。

注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	dic := make([][26]int, n)
	for i, s := range strs {
		for _, c := range s {
			dic[i][int(c-'a')]++
		}
	}

	res := make([][]string, 0)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		temp := []string{strs[i]}
		used[i] = true
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			if dic[i] == dic[j] {
				temp = append(temp, strs[j])
				used[j] = true
			}
		}
		res = append(res, temp)
	}
	return res
}

func groupAnagrams1(strs []string) [][]string {
	//改变哈希表的定义进一步优化
	dic := make(map[[26]int][]string)
	for _, s := range strs {
		var temp [26]int
		for _, c := range s {
			temp[int(c-'a')]++
		}
		if v, ok := dic[temp]; ok {
			v = append(v, s)
			dic[temp] = v
		} else {
			dic[temp] = []string{s}
		}
	}

	var res [][]string
	for _, v := range dic {
		res = append(res, v)
	}
	return res
}
