package codetop

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func EqualStringSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	dic := make(map[string]int)
	for i := range s1 {
		dic[s1[i]]++
	}
	for i := range s1 {
		delete(dic, s2[i])
	}
	return len(dic) == 0
}
