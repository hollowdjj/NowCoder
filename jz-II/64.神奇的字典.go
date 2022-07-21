package jz_II

type MagicDictionary struct {
	t *Trie64
}

/** Initialize your data structure here. */
func Constructor64() MagicDictionary {
	return MagicDictionary{&Trie64{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := range dictionary {
		this.t.add(dictionary[i])
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this.t, searchWord, false)
}

func dfs(node *Trie64, searchWord string, modified bool) bool {
	if len(searchWord) == 0 {
		//长度必须相等
		return node.isEnd && modified
	}
	//深度优先搜索
	index := int(searchWord[0] - 'a')
	if node.nexts[index] != nil && dfs(node.nexts[index], searchWord[1:], modified) {
		return true
	}
	if !modified {
		for i, next := range node.nexts {
			if i != index && next != nil && dfs(next, searchWord[1:], true) {
				return true
			}
		}
	}
	return false
}

type Trie64 struct {
	isEnd bool
	nexts [26]*Trie64
}

func (t *Trie64) add(s string) {
	for _, b := range s {
		index := int(b - 'a')
		if t.nexts[index] == nil {
			t.nexts[index] = &Trie64{}
		}
		t = t.nexts[index]
	}
	t.isEnd = true
}
