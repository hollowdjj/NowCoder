package jz_II

/*
实现一个 MapSum 类，支持两个方法，insert 和 sum：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key，整数表示值 val。
如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

输入：
inputs = ["MapSum", "insert", "sum", "insert", "sum"]
inputs = [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：
[null, null, 3, null, 5]

解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // return 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
*/

type MapSum struct {
	trie *Trie66
}

type Trie66 struct {
	isEnd   bool
	keyVal  int
	pathVal int
	nexts   [26]*Trie66
}

func (t *Trie66) insert(key string, minus, val int) {
	for i := range key {
		index := int(key[i] - 'a')
		if t.nexts[index] == nil {
			t.nexts[index] = &Trie66{keyVal: val, pathVal: val}
		} else {
			t.nexts[index].pathVal += val
			t.nexts[index].pathVal -= minus
		}
		t = t.nexts[index]
	}
	t.isEnd = true
	t.keyVal = val
}

func (t *Trie66) exist(key string) (int, bool) {
	for i := range key {
		index := int(key[i] - 'a')
		if t.nexts[index] == nil {
			return -1, false
		}
		t = t.nexts[index]
	}
	if t.isEnd {
		return t.keyVal, true
	}
	return -1, false
}

/** Initialize your data structure here. */
func Constructor66() MapSum {
	return MapSum{trie: &Trie66{}}
}

func (this *MapSum) Insert(key string, val int) {
	minus := 0
	if v, exist := this.trie.exist(key); exist && v != val {
		minus = v
	}
	this.trie.insert(key, minus, val)
}

func (this *MapSum) Sum(prefix string) int {
	root := this.trie
	for i := range prefix {
		index := int(prefix[i] - 'a')
		if root.nexts[index] == nil {
			return 0
		}
		root = root.nexts[index]
	}
	return root.pathVal
}
