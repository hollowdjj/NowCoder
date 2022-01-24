package tree

/*
给定一棵二叉树，返回齐自底向上的层序遍历。
输入：{1,2,3,4,#,5,6}
返回：[[4,5,6],[2,3],[1]]
*/

/*
vector<vector<int> > levelOrderBottom(TreeNode* root) {
        vector<vector<int>> res;
        stack<vector<int>> s;
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()) {
            auto size = q.size();
            vector<int> temp;
            temp.resize(size);
            for(int i=0;i<size;i++) {
                 auto front = q.front();
                 q.pop();
                temp[i] = front->val;
                if(front->left) q.push(front->left);
                if(front->right) q.push(front->right);
            }
            s.push(std::move(temp));
        }

        while(!s.empty()) {
            res.push_back(std::move(s.top()));
            s.pop();
        }
        return res;
    }
*/
