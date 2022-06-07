package jz_II

/*
某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典
序排列时，返回 true；否则，返回 false。

输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
输出：true
解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。

输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
输出：false
解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"
因为 'l' > '∅'，其中 '∅' 是空白字符，定义为比任何其他字符都小
*/

func isAlienSorted(words []string, order string) bool {
	dic := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		dic[order[i]] = i
	}
	n := len(words)
	for i := 0; i < n-1; i++ {
		l := min(len(words[i]), len(words[i+1]))
		contain := true
		for j := 0; j < l; j++ {
			if dic[words[i][j]] == dic[words[i+1][j]] {
				continue
			} else if dic[words[i][j]] < dic[words[i+1][j]] {
				contain = false
				break
			} else {
				return false
			}
		}
		if contain && len(words[i]) > len(words[i+1]) {
			return false
		}
	}
	return true
}
