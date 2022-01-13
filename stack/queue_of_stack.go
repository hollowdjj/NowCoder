package stack

/*
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。
队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

要求：存储n个元素的空间复杂度为O(n)，插入与删除的时间复杂度都是 O(1)
*/
var stack1 []int
var stack2 []int

func Push(node int) {
	//push的时候，直接往stack1中追加元素即可
	stack1 = append(stack1, node)
}

func Pop() int {
	//首先判断stack2是否为空，若为空，那么将stack1中的元素依次pop给stack2
	if len(stack2) == 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = make([]int, 0)
	}
	res := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}
