package jz

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

此栈包含的方法有：
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素

数据范围：操作数量满足0≤n≤300  ，输入的元素满足∣val∣≤10000
进阶：栈的各个操作的时间复杂度是O(1)，空间复杂度是O(n)
*/

var stack11 []int
var stack22 []int

func Push30(node int) {
	//stack1直接入栈
	stack1 = append(stack11, node)
	if len(stack22) == 0 {
		stack22 = append(stack22, node)
	} else {
		top := stack22[len(stack22)-1]
		if top <= node {
			stack22 = append(stack22, top)
		} else {
			stack22 = append(stack22, node)
		}
	}
}

func Pop30() {
	stack11 = stack1[:len(stack11)-1]
	stack22 = stack2[:len(stack22)-1]
}

func Top() int {
	return stack11[len(stack11)-1]
}
func Min() int {
	return stack22[len(stack22)-1]
}
