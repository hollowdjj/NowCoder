package jz_II

import "sort"

/*
单词数组words的有效编码由任意助记字符串s和下标数组indices组成，且满足：

words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的
子字符串 恰好与 words[i] 相等
给定一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。

输入：words = ["time", "me", "bell"]
输出：10
解释：一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。
words[0] = "time" ，s 开始于 indices[0] = 0 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[1] = "me" ，s 开始于 indices[1] = 2 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[2] = "bell" ，s 开始于 indices[2] = 5 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"

输入：words = ["me","time"]
输出：5
*/

func minimumLengthEncoding(words []string) int {
	res := 0
	t := Trie65{}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	for i := range words {
		if !t.add(words[i]) {
			res += len(words[i]) + 1
		}
	}
	return res
}

type Trie65 struct {
	isEnd bool
	nexts [26]*Trie65
}

//后缀插入，并判断s是否是字典树中某一字符串的后缀
func (t *Trie65) add(s string) bool {
	res := true
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		index := int(s[i] - 'a')
		if t.nexts[index] == nil {
			res = false
			t.nexts[index] = &Trie65{}
		}
		t = t.nexts[index]
	}
	t.isEnd = true
	return res
}
