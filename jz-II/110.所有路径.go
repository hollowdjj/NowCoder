package jz_II

/*
给定一个有n个节点的有向无环图，用二维数组graph表示，请找到所有从0到n-1的路径并输出。
graph 的第i个数组中的单元都表示有向图中i号节点所能到达的下一些结点,若为空，就是没有下一个节点了。

输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
*/

func allPathsSourceTarget(graph [][]int) [][]int {
    var res [][]int
    var dfs func(path []int,i int)
    dfs = func(path []int,i int) {
        if i == len(graph)-1 {
            temp := make([]int,len(path))
            copy(temp,path)
            res = append(res,temp)
            return
        }
        //对第i个节点能够到达的所有节点进行深搜
        for _,node := range graph[i] {
            path = append(path,node)
            dfs(path,node)
            path = path[:len(path)-1]
        }
    }
    dfs([]int{0},0)
    return res
}
