package jz_II

/*
给定两个字符串s1和s2，找到s2中所有s1的变位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
变位词 指字母相同，但排列不同的字符串。

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。
*/

func findAnagrams(s1, s2 string) []int {
	left, right := 0, 0
	dic := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		dic[s1[i]]++
	}
	win := make(map[byte]int)
	count := 0
	res := make([]int, 0)
	for right < len(s2) {
		c := s2[right]
		win[c]++
		if win[c] == dic[c] {
			count++
		}
		for count == len(dic) {
			c = s2[left]
			if win[c] == dic[c] && right-left+1 == len(s1) {
				res = append(res, left)
			}
			win[c]--
			if win[c] < dic[c] {
				count--
			}
			left++
		}
		right++
	}
	return res
}

func findAnagrams1(s1, s2 string) []int {
	//要求s1的排列之一是s2的一个子串，因此可以考虑使用一个长度与s1相等的滑动窗口
	//当滑动窗口中的元素和s1中的元素完全一致时，返回true。由于s1和s2中只有小写
	//字母，因此可以用一个长度为26的数组代替哈希表。
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return []int{}
	}
	dic := [26]int{} //dic[0]表示字母a出现的次数，依次递推
	for _, c := range s1 {
		dic[int(c-'a')]++
	}
	res := make([]int, 0)
	//先创建一个长度为n1的滑动窗口
	win := [26]int{}
	for i := 0; i < n1; i++ {
		win[int(s2[i]-'a')]++
	}
	if win == dic {
		//两数组相同，说明滑动窗口中的元素是s1的一个排列
		res = append(res, 0)
	}
	//滑动窗口往后滑动
	for i := n1; i < n2; i++ {
		win[int(s2[i-n1]-'a')]--
		win[int(s2[i]-'a')]++
		if win == dic {
			res = append(res, i-n1+1)
		}
	}
	return res
}
