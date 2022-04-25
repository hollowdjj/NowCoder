package jz

import (
	"strconv"
	"strings"
)

/*
利用前序遍历进行序列化和反序列化递归版本
*/

func SerializeByPrevOrder(root *TreeNode) string {
	//				1
	//		2		        3
	// 4                          5
	//[1,2,4,#,#,#,3,#,5,#,#]
	res := ""
	var prevOrder func(root *TreeNode)
	prevOrder = func(root *TreeNode) {
		if root == nil {
			res += "#,"
			return
		}
		res += strconv.Itoa(root.Val)
		res += ","
		prevOrder(root.Left)
		prevOrder(root.Right)
	}
	prevOrder(root)
	return res[:len(res)-1]
}

func DeserializeFromPrevOrder(str string) *TreeNode {
	s := strings.Split(str, ",")
	var prevOrder func() *TreeNode
	prevOrder = func() *TreeNode {
		if len(s) == 0 {
			return nil
		}
		if s[0] == "#" {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		s = s[1:]
		root := &TreeNode{Val: val}
		root.Left = prevOrder()
		root.Right = prevOrder()
		return root
	}
	root := prevOrder()
	return root
}

/*
利用后续遍历进行序列化和反序列化
*/
func SerializeByPostOrder(root *TreeNode) string {
	res := ""
	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			res += "#,"
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res += strconv.Itoa(root.Val)
		res += ","
	}
	postorder(root)
	return res[:len(res)-1]
}

func DeserializeFromPostOrder(str string) *TreeNode {
	s := strings.Split(str, ",")
	var postorder func() *TreeNode
	postorder = func() *TreeNode {
		if len(s) == 0 {
			return nil
		}
		if s[len(s)-1] == "#" {
			s = s[:len(s)-1]
			return nil
		}
		//后续遍历的顺序是左右根。因此，当我们从后往前遍历后续遍历数组时的顺序为根右左
		val, _ := strconv.Atoi(s[len(s)-1])
		s = s[:len(s)-1]
		root := &TreeNode{Val: val}
		root.Right = postorder()
		root.Left = postorder()
		return root
	}
	root := postorder()
	return root
}

/*
利用层序遍历进行序列化和反序列化
*/
func SerializeByLevelOrder(root *TreeNode) string {
	res := ""
	q := []*TreeNode{root}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			res += "#,"
		} else {
			res += strconv.Itoa(top.Val)
			res += ","
			q = append(q, top.Left)
			q = append(q, top.Right)
		}
	}
	return res[:len(res)-1]
}

func DeserializeFromLevelOrder(str string) *TreeNode {
	if len(str) == 0 || str[0] == '#' {
		return nil
	}
	s := strings.Split(str, ",")
	val, _ := strconv.Atoi(s[0])
	s = s[1:]
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}

	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if s[0] != "#" {
			val, _ = strconv.Atoi(s[0])
			top.Left = &TreeNode{Val: val}
			q = append(q, top.Left)
		}
		s = s[1:]
		if s[0] != "#" {
			val, _ = strconv.Atoi(s[0])
			top.Right = &TreeNode{Val: val}
			q = append(q, top.Right)
		}
		s = s[1:]
	}
	return root
}
