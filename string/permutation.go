package string

/*
字符串的排列
输入一个长度为 n 字符串，打印出该字符串中字符的所有排列，你可以以任意顺序返回这个字符串数组。
例如输入字符串ABC,则输出由字符A,B,C所能排列出来的所有字符串ABC,ACB,BAC,BCA,CBA和CAB。
*/

var dicPer = map[string]bool{}

func Permutation(str string) []string {
	backTrack([]byte(str), 0)
	res := make([]string, 0)
	for s, _ := range dicPer {
		res = append(res, s)
	}
	return res
}

func backTrack(curr []byte, i int) {
	//递归终止条件。i表示固定str[i]
	if i == len(curr) {
		dicPer[string(curr)] = true
		return
	}
	for j := i; j < len(curr); j++ {
		//交换i,j
		curr[i], curr[j] = curr[j], curr[i]
		backTrack(curr, i+1)
		//不交换i,j
		curr[i], curr[j] = curr[j], curr[i]
	}
}
