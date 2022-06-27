package jz_II

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
*/

//哈希表+滑动窗口
func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	dic := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		dic[s1[i]]++
	}
	win := make(map[byte]int)
	count := 0
	for right < len(s2) {
		c := s2[right]
		win[c]++
		if win[c] == dic[c] {
			count++
		}
		for count == len(dic) {
			c = s2[left]
			if win[c] == dic[c] && right-left+1 == len(s1) {
				return true
			}
			win[c]--
			if win[c] < dic[c] {
				count--
			}
			left++
		}
		right++
	}
	return false
}

//数组+滑动窗口
func checkInclusion1(s1, s2 string) bool {
	//要求s1的排列之一是s2的一个子串，因此可以考虑使用一个长度与s1相等的滑动窗口
	//当滑动窗口中的元素和s1中的元素完全一致时，返回true。由于s1和s2中只有小写
	//字母，因此可以用一个长度为26的数组代替哈希表。
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}
	dic := [26]int{} //dic[0]表示字母a出现的次数，依次递推
	for _, c := range s1 {
		dic[int(c-'a')]++
	}
	//先创建一个长度为n1的滑动窗口
	win := [26]int{}
	for i := 0; i < n1; i++ {
		win[int(s2[i]-'a')]++
	}
	if win == dic {
		//两数组相同，说明滑动窗口中的元素是s1的一个排列
		return true
	}
	//滑动窗口往后滑动
	for i := n1; i < n2; i++ {
		win[int(s2[i-n1]-'a')]--
		win[int(s2[i]-'a')]++
		if win == dic {
			return true
		}
	}
	return false
}
