package jz_II

/*
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现
在单级双链表中。
*/

func flatten(root *Node) *Node {
	res := root
	for root != nil {
		if root.Child != nil {
			next := root.Next

			root.Next = flatten(root.Child)
			root.Next.Prev = root
			root.Child = nil

			for root.Next != nil {
				root = root.Next
			}
			if next != nil {
				next.Prev = root
				root.Next = next
			}

		} else {
			root = root.Next
		}
	}
	return res
}
