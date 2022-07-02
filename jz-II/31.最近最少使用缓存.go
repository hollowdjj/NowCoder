package jz_II

/*
运用所掌握的数据结构，设计和实现一个LRU (Least Recently Used，最近最少使用) 缓存机制 。

实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量capacity初始化LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入
该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为
新的数据值留出空间。
*/

type LRUCache struct {
	dic       map[int]*List
	cap       int
	count     int
	dummyHead *List
	dummyTail *List
}

type List struct {
	Key  int
	Val  int
	Prev *List
	Next *List
}

func Constructor31(capacity int) LRUCache {
	cache := LRUCache{dic: make(map[int]*List), cap: capacity}
	cache.dummyHead, cache.dummyTail = &List{Val: -1}, &List{Val: -1}
	cache.dummyHead.Next = cache.dummyTail
	cache.dummyTail.Prev = cache.dummyHead
	return cache
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.dic[key]
	if ok {
		//将node放到链表头部
		remove(node)
		l.addToHead(node)
		return node.Val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.dic[key]
	if ok {
		node.Val = value
		remove(node)
		l.addToHead(node)
	} else {
		if l.count == l.cap {
			l.removeTail()
			l.count--
		}
		node := &List{Key: key, Val: value}
		l.addToHead(node)
		l.dic[key] = node
		l.count++
	}
}

func (l *LRUCache) addToHead(node *List) {
	next := l.dummyHead.Next
	l.dummyHead.Next = node
	node.Prev = l.dummyHead
	node.Next = next
	if next != nil {
		next.Prev = node
	}
}

func (l *LRUCache) removeTail() {
	node := l.dummyTail.Prev
	if node == nil {
		return
	}
	remove(node)
	delete(l.dic, node.Key)
}

func remove(node *List) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}
