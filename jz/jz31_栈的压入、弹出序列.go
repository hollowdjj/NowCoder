package jz

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。假设压入栈的所有数字均不相等。
例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。
1. 0<=pushV.length == popV.length <=1000
2. -1000<=pushV[i]<=1000
3. pushV 的所有数字均不相同

输入：[1,2,3,4,5],[4,5,3,2,1]
返回值：true
说明：可以通过push(1)=>push(2)=>push(3)=>push(4)=>pop()=>push(5)=>pop()=>pop()=>pop()=>pop()
这样的顺序得到[4,5,3,2,1]这个序列，返回true
*/

func IsPopOrder(push []int, pop []int) bool {
	//使用一个辅助栈模拟
	stack := make([]int, 0)
	j := 0 //指向pop数组
	for i := 0; i < len(push); i++ {
		stack = append(stack, push[i])
		for len(stack) > 0 && stack[len(stack)-1] == pop[j] {
			//如果相等就一直pop
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
