package jz_II

/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word)向前缀树中插入字符串word 。
boolean search(String word)如果字符串word在前缀树中，返回true即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false
*/

type Trie struct {
	isEnd bool
	nexts [26]*Trie
}

/** Initialize your data structure here. */
func Constructor62() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for i, _ := range word {
		index := int(word[i] - 'a')
		if root.nexts[index] == nil {
			root.nexts[index] = &Trie{}
		}
		root = root.nexts[index]
	}
	root.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i, _ := range word {
		index := int(word[i] - 'a')
		if root.nexts[index] == nil {
			return false
		}
		root = root.nexts[index]
	}
	return root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i, _ := range prefix {
		index := int(prefix[i] - 'a')
		if root.nexts[index] == nil {
			return false
		}
		root = root.nexts[index]
	}
	return true
}
