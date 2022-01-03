package array

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有 gas[i] 升油，假设汽车油箱容量无限，从第 i 个
加油站驶往第 (i+1)%n 个加油站需要花费 cost[i] 升油。

请问能否绕环路行驶一周，如果可以则返回出发的加油站编号，如果不能，则返回 -1。
题目数据可以保证最多有一个答案。

*/

func GasStation(gas []int, cost []int) int {
	start := 0 //能够环绕一圈的起点加油站序号
	sum := 0   //绕一圈的剩余油量
	total := 0 //从可能的起点开始往下走的剩余油量

	for i, _ := range gas {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if total < 0 {
			start = i + 1
			total = 0
		}
	}

	if sum < 0 {
		return -1
	}

	return start
}
