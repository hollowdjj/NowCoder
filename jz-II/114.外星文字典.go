package jz_II

/*
现有一种使用英语字母的外星文语言，这门语言的字母顺序与英语顺序不同。给定一个字符串列表words，请你
根据该词典还原出此语言中已知的字母顺序，并按字母递增顺序排列。若不存在合法字母顺序，返回 ""。若存
在多种可能的合法字母顺序，返回其中任意一种 顺序即可。字符串 s 字典顺序小于 字符串 t 有两种情况：
1. 在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于t中字母之前，那么s的字典顺序小于t 。
2. 如果前面min(s.length, t.length)字母都相同，那么s.length < t.length 时，s的字典顺序也小于t。

输入：words = ["wrt","wrf","er","ett","rftt"]
输出："wertf"

输入：words = ["z","x"]
输出："zx"
*/

func alienOrder(words []string) string {
	//建图
	graph := [26][]byte{}
	indeg := [26]int{}
	for i := range indeg {
		indeg[i] = -1 //便于判断哪些字符出现过，哪些没有出现过
	}
	n := len(words)
	for i := range words[0] {
		indeg[int(words[0][i]-'a')] = 0
	}
loop:
	for i := 1; i < n; i++ {
		s1, s2 := words[i-1], words[i]
		for j := range words[i] {
			if index := int(words[i][j] - 'a'); indeg[index] == -1 {
				indeg[index] = 0
			}
		}
		for j := 0; j < len(s1) && j < len(s2); j++ {
			if s1[j] != s2[j] {
				index1 := int(s1[j] - 'a')
				index2 := int(s2[j] - 'a')
				graph[index1] = append(graph[index1], s2[j])
				indeg[index2]++
				continue loop
			}
		}
		//对于abc ab的情况
		if len(s1) > len(s2) {
			return ""
		}
	}
	//BFS实现拓扑排序
	//将所有入度为0的节点入队
	var res []byte
	var q []int
	for i, deg := range indeg {
		if deg == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		res = append(res, 'a'+byte(front))
		for _, node := range graph[front] {
			i := int(node - 'a')
			if indeg[i]--; indeg[i] == 0 {
				q = append(q, i)
			}
		}
	}
	//统计字符数量unique
	count := 0
	for i := range indeg {
		if indeg[i] >= 0 {
			count++
		}
	}
	if len(res) == count {
		return string(res)
	}
	return ""
}
