package tree

import "nowcoder/utility"

/*
合并二叉树
已知两颗二叉树，将它们合并成一颗二叉树。
合并规则是：都存在的结点，就将结点值加起来，否则空的位置就由另一个树的结点来代替
*/

func MergeTrees(t1 *utility.TreeNode, t2 *utility.TreeNode) *utility.TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)
	return t1
}
