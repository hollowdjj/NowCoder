package tree

/*
输出二叉树的右视图
请根据二叉树的前序遍历，中序遍历恢复二叉树，并打印出二叉树的右视图

数据范围： 0≤n≤10000
要求： 空间复杂度 O(n)，时间复杂度O(n)

如输入[1,2,4,5,3],[4,2,5,1,3]时，通过前序遍历的结果[1,2,4,5,3]和中序遍历的结果
可以重建出二叉树：
						1
					2        3
                4       5
右视图为：[1,3,5]
*/

/*
vector<int> solve(vector<int>& xianxu, vector<int>& zhongxu) {
        //将中序遍历数组中的元素与其对应的索引存储在map中
        map<int,int> dic;
        for(int i=0;i<zhongxu.size();i++)
        {
            dic[zhongxu[i]] = i;
        }
        //重建二叉树
        auto root = reconstruct(xianxu,0,xianxu.size()-1,zhongxu,0,zhongxu.size()-1,dic);
        //层序遍历二叉树，但只保存每一层的最后一个数字
        vector<int> res;
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty())
        {
            auto size = q.size();
            res.push_back(q.back()->val);
            while(size--)
            {
                auto front = q.front();
                q.pop();
                if(front->left)  q.push(front->left);
                if(front->right) q.push(front->right);
            }
        }
        return res;
    }

    TreeNode*  reconstruct(vector<int>& pre,int pre_left,int pre_right,
                          vector<int>& in,int in_left,int in_right,const map<int,int>& dic)
    {
        //递归终止条件
        if (pre_left > pre_right || in_left > in_right) {
            return nullptr;
        }
        int index = dic.find(pre[pre_left])->second; //root在中序数组中的索引
        TreeNode* root = new TreeNode(pre[pre_left]);
        int count = index - in_left; //root左子树的节点数量
        root->left = reconstruct(pre,pre_left+1,pre_left+count,in,in_left,index-1,dic);
        root->right = reconstruct(pre,pre_left+count+1,pre_right,in,index+1,in_right,dic);
        return root;
    }
*/

func RightView(root *TreeNode) []int {
	//要获取二叉树的右视图，那么我们就需要一直走右子节点
	res := make([]int, 0)
	var dfs func(root *TreeNode, i int)
	dfs = func(root *TreeNode, i int) {
		if root == nil || i > len(res) {
			return
		}
		//类似于前序遍历，但是一直走右子节点。
		if i == len(res) {
			res = append(res, root.Val)
		}
		dfs(root.Right, i+1)
		dfs(root.Left, i+1)
	}
	dfs(root, 0)
	return res
}
