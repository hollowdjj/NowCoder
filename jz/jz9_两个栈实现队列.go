package jz

/*
用两个栈来实现一个队列
*/

var stack1 []int
var stack2 []int

func Push(node int) {
	//直接添加到stack1中
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		//将stack1中的元素全部pop到stack2中
		n := len(stack1)
		for i := n - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}
	ret := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return ret
}
