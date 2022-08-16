package jz_II

/*
在字典（单词列表） wordList 中，从单词 beginWord 和 endWord 的转换序列 是一个按下述规格形成的序列：

序列中第一个单词是 beginWord。
序列中最后一个单词是 endWord。
每次转换只能改变一个字母。
转换过程中的中间单词必须是字典 wordList 中的单词。
给定两个长度相同但内容不同的单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的
最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回0。
*/

/*

建图然后BFS即可

class Solution {
public:
    bool canTransform(const string& s1,const string& s2) {
        if(s1.size() != s2.size()) {
            return false;
        }
        int n = s1.size();
        int diff = 0;
        for(int i=0;i<n;i++) {
            if(s1[i] != s2[i]) {
                diff++;
                if(diff > 1) {
                    return false;
                }
            }
        }
        return true;
    }

    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        auto id = map<string,int>();
        for(auto& word : wordList) {
            id[word] = id.size();
            cout<<word<<" "<<id[word]<<endl;
        }
        //若endword不在wordlist中，直接返回0
        if(id.find(endWord) == id.end()) {
            return 0;
        }
        if(id.find(beginWord) == id.end()) {
            id[beginWord] = id.size();
        }
        cout<<beginWord<<" "<<id[beginWord]<<endl;
        wordList.push_back(beginWord);
        //初始化邻接矩阵
        int n = id.size();
        vector<vector<int>> grid(n,vector<int>(n,0));
        //建图，若能转换则用一条边相连
        for(int i=0;i<n;i++) {
            for (int j=0;j<n;j++) {
                if(i == j || canTransform(wordList[i],wordList[j])) {
                    grid[i][j] = 1;
                }
                cout<<grid[i][j]<<" ";
            }
            cout <<endl;
        }
        //bfs
        queue<int> q;
        q.push(id[beginWord]);
        int step = 1;
        int target = id[endWord];
        vector<int> visit(n,0);
        visit[id[beginWord]] = 1;
        while(!q.empty()) {
            step++;
            int size = q.size();
            for(int i=0;i<size;i++) {
                auto front = q.front();
                cout<<front<<" "<<step<<endl;
                q.pop();
                for(int j=0;j<n;j++) {
                    if (grid[front][j] == 0) {
                        continue;
                    }
                    if (!visit[j]) {
                        if(j == target) {
                            return step;
                        }
                        q.push(j);
                        visit[j] = 1;
                    }
                }
            }
        }
        return 0;
    }
};
*/
