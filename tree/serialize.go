package tree

/*
序列化二叉树
*/

/*
class Solution {
public:
    char* Serialize(TreeNode *root) {
        if(!root){
            return nullptr;
        }
        queue<TreeNode*> q;
        q.push(root);
        vector<string> levels;
        while(!q.empty()) {
            auto size = q.size();
            string temp;
            for(int i=0;i<size;i++) {
                auto front = q.front();
                q.pop();
                if(!front) {
                    temp += "#,";
                } else {
                    temp += (std::to_string(front->val)+",");
                    q.push(front->left);
                    q.push(front->right);
                }
            }
            levels.push_back(std::move(temp));
        }
        levels.pop_back();
        string res;
        for(auto& str : levels) {
            res += str;
        }
        char *ret = new char[res.length() + 1];
        strcpy(ret, res.c_str());
       return ret;
    }

    TreeNode* Deserialize(char *str) {
        if (str == nullptr) {
            return nullptr;
        }
        // 可用string成员函数
        string s(str);
        if (str[0] == '#') {
            return nullptr;
        }

        // 构造头结点
        queue<TreeNode*> nodes;
        TreeNode *ret = new TreeNode(atoi(s.c_str()));
        s = s.substr(s.find_first_of(',') + 1);
        nodes.push(ret);
        // 根据序列化字符串再层次遍历一遍，来构造树
        while (!nodes.empty() && !s.empty())
        {
            TreeNode *node = nodes.front();
            nodes.pop();
            if (s[0] == '#')
            {
                node->left = nullptr;
                s = s.substr(2);
            }
            else
            {
                node->left = new TreeNode(atoi(s.c_str()));
                nodes.push(node->left);
                s = s.substr(s.find_first_of(',') + 1);
            }

            if (s[0] == '#')
            {
                node->right = nullptr;
                s = s.substr(2);
            }
            else
            {
                node->right = new TreeNode(atoi(s.c_str()));
                nodes.push(node->right);
                s = s.substr(s.find_first_of(',') + 1);
            }
        }
        return ret;
    }
};
*/
