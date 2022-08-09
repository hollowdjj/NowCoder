package jz_II

import (
	"fmt"
	"strconv"
	"strings"
)

/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，只需要保证一
个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
*/

//前序遍历序列化二叉树
func serializeByPreOrder(root *TreeNode) string {
	res := ""
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			res += "#,"
			return
		}
		res += fmt.Sprintf("%d,", root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res[:len(res)-1]
}

//反序列化前序遍历数组
func deserializePreOrder(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	i := 0
	var preorder func() *TreeNode
	preorder = func() *TreeNode {
		if i == len(nodes) || nodes[i] == "#" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(nodes[i])
		root := &TreeNode{Val: val}
		i++
		root.Left = preorder()
		root.Right = preorder()
		return root
	}
	return preorder()
}

//层序遍历序列化二叉树
func serializeByLevelOrder(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	q := []*TreeNode{root}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		if front == nil {
			res = append(res, "#")
		} else {
			res = append(res, strconv.Itoa(front.Val))
			q = append(q, front.Left)
			q = append(q, front.Right)
		}
	}
	return strings.Join(res, ",")
}

func deserializeLevelOrder(data string) *TreeNode {
	strs := strings.Split(data, ",")
	if data=="" || strs[0] == "#" {
		return nil
	}
	rootVal,_ := strconv.Atoi(strs[0])
	res := &TreeNode{Val:rootVal}
	q := []*TreeNode{res}
	i := 1
	for i < len(strs) {
		root := q[0]
        q = q[1:]
		if strs[i] != "#" {
			leftVal,_ := strconv.Atoi(strs[i])
			root.Left = &TreeNode{Val:leftVal}
			q = append(q,root.Left)
		}
		i++
		if strs[i] != "#" {
			rightVal,_ := strconv.Atoi(strs[i])
			root.Right = &TreeNode{Val:rightVal}
			q = append(q,root.Right)
		}
		i++
	}
	return res
}
