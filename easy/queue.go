package easy

import "fmt"

/*
两个栈来实现一个队列，使用n个元素来完成n次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。
队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

要求：存储n个元素的空间复杂度为 O(n)，插入与删除的时间复杂度都是 O(1)
*/
var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack1) == 0 {
		return 0
	}
	res := stack1[0]
	stack1 = stack1[1:]
	return res
}

/*
需要注意的是，在Go中由于没有提供stack，所以使用一个数组就可以解题。最正确的做法应该是需要去模拟栈先入后
出的性质，然后再解题。
*/
type stack []int    //定义一个stack类型
type Queue struct { //两个stack构成的队列
	stack1 stack
	stack2 stack
}

func (s *stack) Pop() int {
	//先入后出
	if len([]int(*s)) == 0 {
		panic("Empty queue")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return res
}

func (s *stack) Push(node int) {
	*s = append(*s, node)
}

func (q *Queue) Push(node int) {
	q.stack1.Push(node) //stack1用于入队
}

func (q *Queue) Pop() int {
	//队列是先入先出，后入后出
	if len(q.stack2) == 0 {
		//将stack1中的元素逆序放入stack2中
		for len(q.stack1) != 0 {
			q.stack2.Push(q.stack1.Pop())
		}
		//清空stack1
		q.stack1 = q.stack1[:0]
	}

	return q.stack2.Pop()
}

func TestQueue() {
	var queue Queue

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	//fmt.Println(queue.Pop())
}
