package tree

//TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//SliceToBinaryTree 将一个表示二叉树层序遍历结果的slice，即arr[index:]，转换成一颗二叉树。
//二叉树必须是一颗完全二叉树，即所有不存的节点都需要用#表示
func SliceToBinaryTree(arr []int, index int) *TreeNode {
	if index >= len(arr) || arr[index] == '#' {
		return nil
	}
	//根节点
	root := &TreeNode{Val: arr[index]}
	//得到左子节点
	root.Left = SliceToBinaryTree(arr, 2*index+1)
	//得到右子节点
	root.Right = SliceToBinaryTree(arr, 2*index+2)

	return root
}

//LevelOrder 层序遍历打印二叉树，不存在的节点使用#代替。
//func (t *TreeNode) LevelOrder() (res []int) {
//	//创建一个队列
//	q := CreateQueue()
//	//头结点入队
//	q.Push(t)
//
//	for !q.Empty() {
//		//首元素出队
//		node := (*q.Pop()).(*TreeNode)
//		if node == nil {
//			res = append(res, '#')
//		} else {
//			res = append(res, node.Val)
//			if !(node.Left == nil && node.Right == nil) {
//				//左子节点入队
//				q.Push(node.Left)
//				//右子节点入队
//				q.Push(node.Right)
//			}
//		}
//	}
//
//	return
//}
