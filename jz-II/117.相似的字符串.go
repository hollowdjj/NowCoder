package jz_II

/*
如果交换字符串X中的两个不同位置的字母，使得它和字符串Y相等，那么称X和Y两个字符串相似。如果这两个字符串本身是相等的，那它们
也是相似的。

例如，"tars"和"rats"是相似的(交换0与2的位置)；"rats" 和 "arts" 也是相似的，但是"star"不与"tars"，"rats"，或 "arts" 相似。
总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。注意，"tars" 和 "arts" 是在同一组中，即使它们
并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给定一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个 字母异位词 。请问 strs 中有多少个相似字符串组？

字母异位词（anagram），一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。
*/

/*
本质上就是求图的联通量而已：
class Solution {
public:
    bool similar(const string& s1,const string& s2) {
        int diff = 0;
        int n = s1.size();
        for(int i=0;i<n;i++) {
            if(s1[i] != s2[i]) {
                diff++;
            }
            if(diff > 2) {
                return false;
            }
        }
        return true;
    }

    int numSimilarGroups(vector<string>& strs) {
        //建图
        int n = strs.size();
        vector<vector<int>> graph(n);
        for(int i=0;i<n;i++) {
            for(int j=0;j<n;j++) {
                if(i!=j && similar(strs[i],strs[j])) {
                    graph[i].push_back(j);
                }
            }
        }

        //计算联通量
        vector<int> visit(n,0);
        int res = 0;
        auto bfs = [&graph,&visit](int start) {
            queue<int> q;
            q.push(start);
            visit[start] = 1;
            while(!q.empty()) {
                auto front = q.front();
                q.pop();
                for(auto& j : graph[front]) {
                    if(!visit[j]) {
                        q.push(j);
                        visit[j] = 1;
                    }
                }
            }
        };
        for(int i=0;i<n;i++) {
            if(!visit[i]) {
                bfs(i);
                res++;
            }
        }
        return res;
    }
};
*/
