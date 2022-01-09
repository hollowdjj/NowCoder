package stack

import "nowcoder/utility"

/*
给定一棵二叉树，分别按照二叉树先序，中序和后序打印所有的节点。
数据范围：0≤n≤1000，树上每个节点的val值满足 0≤val≤100
要求：空间复杂度 O(n)，时间复杂度 O(n)

				 1
              /    \
             2      3
           /   \  /   \
          4    5  6    7
前序遍历(根左右)：[1,2,4,5,3,6,7]
中序遍历(左根右)：[4,2,5,1,6,7,3]
后序遍历(左右根)：[4,5,2,6,7,3,1]

此题采用递归的话，时间复杂度为O(n)，空间复杂度为O(1)。除此之外，我们还可以使用栈来实现
一个非递归的解法。
*/

func ThreeOrders(root *utility.TreeNode) [][]int {
	res := make([][]int, 3)
	res[0] = preOrder(root)
	res[1] = inOrder(root)
	res[2] = postOrder(root)
	return res
}

func preOrder(root *utility.TreeNode) []int {
	var res []int
	var do func(node *utility.TreeNode)
	do = func(node *utility.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		do(node.Left)
		do(node.Right)
	}
	do(root)
	return res
}

func preOrderWithStack(root *utility.TreeNode) []int {
	//前序遍历中，节点的顺序为：根节点-->左子节点-->右子节点。因此，节点入栈的顺序
	//应该是根节点-->右子节点-->左子节点
	var res []int
	var stack utility.Stack
	stack.Push(root)
	for !stack.Empty() {
		top := (*stack.Pop()).(*utility.TreeNode)
		res = append(res, top.Val)
		if top.Right != nil {
			stack.Push(top.Right)
		}
		if top.Left != nil {
			stack.Push(top.Left)
		}
	}
	return res
}

func inOrder(root *utility.TreeNode) []int {
	var res []int
	var do func(node *utility.TreeNode)
	do = func(node *utility.TreeNode) {
		if node == nil {
			return
		}
		do(node.Left)
		res = append(res, node.Val)
		do(node.Right)
	}
	do(root)
	return res
}

func inOrderWithStack(root *utility.TreeNode) []int {
	//中序遍历中，节点的顺序为左子节点-->根节点-->右子节点。这里使用栈的话要复杂一些。
	//1. 访问以当前结点p为根结点的[最左孩子]n;
	//2. 访问n，并将p更新至其右子树上：p = p->right
	//3. 若p无右子树，则说明p是某颗子树的叶子结点，则在下一次循环中访问p；若p有右子树，则按照上述相同步骤访问该右子树，最后访问p。
	//                     1
	//				    2     3
	//				 4    5
	//首先，1,2,4依次入栈。然后访问栈顶元素，也就是4。4没有右子树，那么就接着访问当前的栈顶元素2。2存在右子树，那么就按照步骤1,2
	//重复。
	var res []int
	var stack utility.Stack
	for root != nil || !stack.Empty() {
		for root != nil {
			//寻找最左的孩子
			stack.Push(root)
			root = root.Left
		}
		top := (*stack.Pop()).(*utility.TreeNode) //访问
		res = append(res, top.Val)
		root = root.Right //右子树
	}
	return res
}

func postOrder(root *utility.TreeNode) []int {
	var res []int
	var do func(node *utility.TreeNode)
	do = func(node *utility.TreeNode) {
		if node == nil {
			return
		}
		do(node.Left)
		do(node.Right)
		res = append(res, node.Val)
	}
	do(root)
	return res
}

func postOrderWithStack(root *utility.TreeNode) []int {
	//1. 定义last指针，用以代表[上次访问的位置]，p结点用来遍历
	//2. 由于后续遍历需要先访问左子树和右子树、最后再访问根结点，因此在满足：1) 该结点没有右子树；
	//   2)该结点有右子树，且上次访问的结点(last)指针为其右子树，即该右子树已经被访问过了才能访问该结点
	//3. 在不满足上述条件的情况下，说明当前结点还不能被访问，故先访问其右子树:p = p.right
	//4. 将p结点置为NULL，防止下次while循环时重复访问。
	//                     1
	//				    2     3
	//				 4     5
	var res []int
	var stack utility.Stack
	var last *utility.TreeNode
	p := root
	for !stack.Empty() || p != nil {
		//找到最左边的节点
		for p != nil {
			stack.Push(p)
			p = p.Left
		}
		p = (*stack.Top()).(*utility.TreeNode)
		if p.Right == nil || last == last.Right {
			stack.Pop()
			res = append(res, p.Val)
			last = p
			p = nil
		} else {
			p = p.Right
		}
	}
	return res
}
