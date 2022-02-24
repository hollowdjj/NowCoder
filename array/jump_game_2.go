package array

/*
几步可以从头跳到尾

给你一个长度为 n 的数组 a。
ai 表示从 i 这个位置开始最多能往后跳多少格。
求从 1 开始最少需要跳几次就能到达第 n 个格子。
*/

func Jump(n int, A []int) int {
	prev := 0 //前一个可以跳到的最远位置
	curr := 0 //当前能跳到的最远位置。即A[0..i]所有跳法中能跳到的最远位置
	res := 0
	//i不能到最后一格。因为题目要求的是跳到最后一格
	for i := 0; i < n-1; i++ {
		curr = max(curr, i+A[i])
		if i == prev {
			//如果走到了上一次能到达的最远位置，跳跃次数加1。意思就是说我可以从之前
			//某个地方一下跳到当前位置。
			res++
			prev = curr
		}
	}
	return res
}
