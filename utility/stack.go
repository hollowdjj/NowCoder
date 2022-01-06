package utility

//Element 栈中的元素
type Element interface{}

//Stack 栈
type Stack struct {
	elems []Element
}

type Stacker interface {
	Push(elem Element)
	Pop() *Element
	Top() *Element
	Empty() bool
}

func (s *Stack) Push(elem Element) {
	s.elems = append(s.elems, elem)
}

func (s *Stack) Pop() *Element {
	if s.Empty() {
		panic("Pop on empty stack")
	}
	res := s.Top()
	s.elems = (s.elems)[:len(s.elems)-1]
	return res
}

func (s *Stack) Top() *Element {
	if s.Empty() {
		panic("No elements in stack")
	}
	return &(s.elems)[len(s.elems)-1]
}

func (s *Stack) Empty() bool {
	if len(s.elems) == 0 {
		return true
	}
	return false
}
