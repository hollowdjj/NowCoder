package stack

import "nowcoder/utility"

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

此栈包含的方法有：
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素
*/

/*
思路在于使用两个栈，其中一个栈执行正常的操作。第二个栈用于存放最小值。关键点在于，
当要入栈的元素v大于minValue栈的栈顶元素时，minValue入栈的不是v而是其栈顶元素。
*/
var normal utility.Stack
var minValue utility.Stack

func Push1(node int) {
	normal.Push(node)
	if minValue.Empty() {
		minValue.Push(node)
	} else if top := (*minValue.Top()).(int); top <= node {
		minValue.Push(top)
	} else {
		minValue.Push(node)
	}
}
func Pop1() {
	normal.Pop()
	minValue.Pop()
}
func Top() int {
	return (*normal.Top()).(int)
}
func Min() int {
	return (*minValue.Top()).(int)
}
