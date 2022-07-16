package jz_II

/*
在英语中，有一个叫做 词根(root) 的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词
为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

现在，给定一个由许多词根组成的词典和一个句子，需要将句子中的所有继承词用词根替换掉。如果继承词有许
多可以形成它的词根，则用最短的词根替换它。需要输出替换之后的句子。

输入：dictionary=["cat","bat","rat"],sentence="the cattle was rattled by the battery"
输出："the cat was rat by the bat"
*/

func replaceWords(dictionary []string, sentence string) string {
	root := &TrieNode{}
	for i, _ := range dictionary {
		add53(root, dictionary[i])
	}
	res := ""
	n := len(sentence)
	left, right := 0, 0
	for right <= n {
		if right == n || sentence[right] == ' ' {
			str := sentence[left:right]
			replace := find63(root, str)
			if replace == "" {
				res += str
			} else {
				res += replace
			}
			res += " "
			left = right + 1
		}
		right++
	}
	return res[:len(res)-1]
}

//前缀树节点
type TrieNode struct {
	word  string //以该节点结尾的字符串
	nexts [26]*TrieNode
}

func add53(root *TrieNode, s string) {
	if root == nil {
		return
	}
	for i, _ := range s {
		index := int(s[i] - 'a')
		if root.nexts[index] == nil {
			root.nexts[index] = &TrieNode{}
		}
		root = root.nexts[index]
		if root.word != "" {
			return
		}
	}
	root.word = s
}

func find63(root *TrieNode, s string) string {
	if root == nil {
		return ""
	}
	for i, _ := range s {
		index := int(s[i] - 'a')
		if root.nexts[index] == nil {
			return ""
		}
		root = root.nexts[index]
		if root.word != "" {
			break
		}
	}
	return root.word
}
