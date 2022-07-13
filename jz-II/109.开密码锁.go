package jz_II

/*
一个密码锁由4个环形拨轮组成，每个拨轮都有10个数字：'0','1','2','3','4','5','6','7','8','9'。
每个拨轮可以自由旋转：例如把'9'变为'0'，'0'变为'9'。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为'0000'，一个代表四个拨轮的数字的字符串。列表deadends包含了一组死亡数字，一旦拨轮的
数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，请给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
*/

func openLock(deadends []string, target string) int {
	//将死亡数字放入哈希表中
	dead := make(map[string]bool)
	for _, s := range deadends {
		dead[s] = true
	}
	if _, ok := dead[target]; ok {
		return -1
	}
	if _, ok := dead["0000"]; ok {
		return -1
	}
	q := []string{"0000"}
	visit := map[string]bool{"0000": true}
	res := 0
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:]
			if _, ok := dead[front]; ok {
				continue
			}
			if front == target {
				return res
			}
			//遍历所有可能
			for i := 0; i < 4; i++ {
				up := upOne(front, i)
				down := downOne(front, i)
				if _, ok := visit[up]; !ok {
					q = append(q, up)
					visit[up] = true
				}
				if _, ok := visit[down]; !ok {
					q = append(q, down)
					visit[down] = true
				}
			}
		}
		res++
	}
	return -1
}

func upOne(s string, i int) string {
	res := []byte(s)
	if res[i] == '9' {
		res[i] = '0'
	} else {
		res[i]++
	}
	return string(res)
}

func downOne(s string, i int) string {
	res := []byte(s)
	if res[i] == '0' {
		res[i] = '9'
	} else {
		res[i]--
	}
	return string(res)
}
