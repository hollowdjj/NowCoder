package utility

/*
QueueInt一个元素类型为int的队列
*/

type Item interface {
}

type Queue struct {
	items []Item
}

type Queuer interface {
	Enqueue(i Item)
	Dequeue() *Item
	IsEmpty() bool
	Size() int
}

//Enqueue 将元素插入到队列尾部。注意，这里的参数不能设置为*Item类型
//因为，空的interface的动态类型是在赋值的时候隐式确定的。所以，这里
//必须是值类型，从而才能确定Item的动态类型。
func (q *Queue) Enqueue(i Item) {
	q.items = append(q.items, i)
}

//Dequeue 首元素出队
func (q *Queue) Dequeue() *Item {
	res := &q.items[0]
	q.items = q.items[1:]
	return res
}

//IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	n := len(q.items)
	if n > 0 {
		return false
	}

	return true
}

func (q *Queue) Size() int {
	return len(q.items)
}

//CreateQueue 创建一个空的队列
func CreateQueue() *Queue {
	emptyQueue := new(Queue)
	return emptyQueue
}
