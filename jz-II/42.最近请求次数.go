package jz_II

/*
写一个RecentCounter类来计算特定时间范围内最近的请求。
请实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去
3000毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证每次对 ping 的调用都使用比之前更大的 t 值。
*/

type RecentCounter struct {
	q []int
}

func Constructor42() RecentCounter {
	r := RecentCounter{q: make([]int, 0)}
	return r
}

func (r *RecentCounter) Ping(t int) int {
	r.q = append(r.q, t)
	for r.q[0] < t-3000 {
		r.q = r.q[1:]
	}
	return len(r.q)
}
