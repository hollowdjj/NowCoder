package utility

//Stack 栈
type Stack struct {
	elems []interface{}
}

type Stacker interface {
	Push(elem interface{})
	Pop() *interface{}
	Top() *interface{}
	Empty() bool
}

func (s *Stack) Push(elem interface{}) {
	s.elems = append(s.elems, elem)
}

func (s *Stack) Pop() *interface{} {
	if s.Empty() {
		panic("Pop on empty stack")
	}
	res := s.Top()
	s.elems = (s.elems)[:len(s.elems)-1]
	return res
}

func (s *Stack) Top() *interface{} {
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
