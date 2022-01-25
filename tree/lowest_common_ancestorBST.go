package tree

/*
二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，最近公共祖先LCA(T,p,q)表示一个节点x，
满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值； 若它的右子树不空，
则右子树上所有节点的值均大于它的根节点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。
数据范围:
3<=节点总数<=10000
0<=节点值<=10000
*/

func LowestCommonAncestorBST(root *TreeNode, p, q int) int {
	for {
		//如果p和q都大于root的值，说明p和q都在root的右边
		if p > root.Val && q > root.Val {
			root = root.Right
		} else if p < root.Val && q < root.Val {
			root = root.Left
		} else {
			//如何p和q在root的左右两边，那么root就是它们的最近公共祖先
			break
		}
	}
	return root.Val
}
