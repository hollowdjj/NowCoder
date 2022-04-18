package jz

/*
输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针random指向一个随机节点）
请对此链表进行深拷贝，并返回拷贝后的头结点。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）。
下图是一个含有5个结点的复杂链表。图中实线箭头表示next指针，虚线箭头表示random指针。为简单起见，指向null的指针没有画出。
*/

func Clone(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	//这道题的关键在于是要执行深拷贝，因此不太好处理random指针。一个可行的方法是，先拷贝链表
	//然后用哈希表记录一下每个节点相对应的random节点。时间复杂度和空间复杂度都是O(n)。一种空
	//间复杂度为O(1)的方法如下：

	//1. 首先深拷贝链表，形成: 原1-->拷1-->原2-->拷2--> .... -->原n-->拷n-->null 的结构
	nHead := head
	for nHead != nil {
		temp := &RandomListNode{Label: nHead.Label}
		temp.Next = nHead.Next
		nHead.Next = temp
		nHead = temp.Next
	}

	//2. 然后处理random指针。由于形成了1中所示的结构，对于原链表的某一结点的random指针可知，其深拷贝对应
	//的结点的random指针为random.Next
	old, clone := head, head.Next
	for old != nil {
		random := old.Random
		if random != nil {
			clone.Random = random.Next
		}
		if old.Next != nil {
			old = old.Next.Next
		}
		if clone.Next != nil {
			clone = clone.Next.Next
		}
	}

	//3.拆分链表
	old, clone = head, head.Next
	ret := clone
	for old != nil {
		if old.Next != nil {
			old.Next = old.Next.Next
		}
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}
		old = old.Next
		clone = clone.Next
	}
	return ret
}
