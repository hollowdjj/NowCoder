package nowcoder

type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	for pHead != nil {
		next := pHead.Next    //先把后一个结点保存了
		pHead.Next = prev	  //修改当前结点后后续结点为当前结点的前一个结点
		prev = pHead		  //update
		pHead = next		  //update
	}
	return prev
}

func TestReverseList() {
	pHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := pHead
	for i:=1;i<5;i++ {
		newNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		temp.Next = newNode
		temp = newNode
	}
	ReverseList(pHead)
}
