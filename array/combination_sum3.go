package array

/*
找出所有相加之和是 n 的 k 个数的组合。组合中只含有 0~9的正整数，且保证每种组合中不含有相同的数字。
保证一定有解。结果按字典序升序输出。

输入：3,7
输出：[[1,2,4]]
*/

func Combination3(k int, n int) [][]int {
	var res [][]int
	var bt func(i int, target int, temp []int)
	bt = func(i int, target int, temp []int) {
		if len(temp) > k {
			return
		}
		if target == 0 && len(temp) == k {
			res = append(res, temp)
		}
		for j := i; j <= 9; j++ {
			temp = append(temp, j)
			bt(j+1, target-j, temp)
			//回溯
			t := make([]int, len(temp)-1)
			copy(t, temp[:len(temp)-1])
			temp = t
		}
	}
	bt(1, n, []int{})
	return res
}

func backTrack3() {

}
