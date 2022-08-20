package codetop

type Trie struct {
	nexts [26]*Trie
	isEnd bool
	val   string
}

func add(root *Trie, str string) {
	if root == nil {
		return
	}
	for i := range str {
		index := int(str[i] - 'a')
		if root.nexts[index] == nil {
			root.nexts[index] = &Trie{}
		}
		root = root.nexts[index]
	}
	root.isEnd = true
	root.val = str
}

func search(root *Trie, target string) []string {
	if root == nil {
		return nil
	}
	for i := range target {
		index := int(target[i] - 'a')
		if root.nexts[index] == nil {
			return nil
		}
		root = root.nexts[index]
	}
	//使用广度优先搜索
	var res []string
	q := []*Trie{root}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		if front.isEnd {
			res = append(res, front.val)
		}
		for _, node := range front.nexts {
			if node != nil {
				q = append(q, node)
			}
		}
	}
	return res
}

func CommonPrefix(strs []string, prefix string) []string {
	root := &Trie{}
	for i := range strs {
		add(root, strs[i])
	}
	return search(root, prefix)
}
