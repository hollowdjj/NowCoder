package jz_II

/*
给定一个二叉搜索树的根节点root和一个整数k, 请判断该二叉搜索树中是否存在两个节点它们的值之和等于k。
假设二叉搜索树中节点的值均唯一。
*/

func findTarget1(root *TreeNode, k int) bool {
	//中序遍历+哈希。时间复杂度O(N),空间复杂度O(N)
	dic := make(map[int]bool)
	res := false
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if res || root == nil {
			return
		}
		inorder(root.Left)
		if dic[k-root.Val] {
			res = true
			return
		}
		dic[root.Val] = true
		inorder(root.Right)
	}
	inorder(root)
	return res
}

func findTarget(root *TreeNode, k int) bool {
	//中序遍历数组+双指针。时间复杂度O(N),空间复杂度O(N)
	arr := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		arr = append(arr, root.Val)
		inorder(root.Right)
	}
	inorder(root)

	left, right := 0, len(arr)-1
	for left < right {
		t := arr[left] + arr[right]
		if t < k {
			left++
		} else if t > k {
			right--
		} else {
			return true
		}
	}
	return false
}
