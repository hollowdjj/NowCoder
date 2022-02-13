package string

/*
给表达式添加运算符
给定一个仅包含数字的字符串 num 和一个目标值 target，在 num 的数字之间添加二元运算符
"+" , "-" 或 "*" ,返回所有能够得到目标值的表达式。

数据范围：字符串长度满足1≤n≤10，nums中仅包含数字，∣target∣≤10^9

例如："123",6
返回值：["1+2+3","1*2*3"]
*/

/*
set<string> res;

vector<string> addOpt(string num, int target) {
dfs(num,1,string{num[0]},target);
vector<string> ret;
for(auto& s : res) {
ret.push_back(s);
}
return ret;
}

void dfs(string num,int i,string curr,int target) {
//递归终止条件
auto n = num.size();
if(i == n) {
if(cal(curr) == target) {
res.insert(curr);
}
return;
}

auto temp = curr;
for(int j=i;j<n;j++) {
dfs(num,j+1,curr+"+"+string{num[j]},target);
dfs(num,j+1,curr+"-"+string{num[j]},target);
dfs(num,j+1,curr+"*"+string{num[j]},target);
}
}

int cal(string opt) {
//计算表达式的值
stack<int> numbers;
int val = 0;
auto sign = '+';
for(int i=0;i<opt.size();i++) {
if(opt[i] <= '9' && opt[i]>= '0') {
val = opt[i]-'0';
}
if(i == opt.size()-1 || opt[i]=='+'||opt[i]=='-'||opt[i]=='*') {
if(sign == '+') {
numbers.push(val);
} else if(sign == '-') {
numbers.push(-1*val);
} else if (sign == '*') {
auto top = numbers.top();
numbers.pop();
numbers.push(top*val);
}
sign = opt[i];
val = 0;
}
}
int res = 0;
while(!numbers.empty()) {
res += numbers.top();
numbers.pop();
}
return res;
}
};
*/
