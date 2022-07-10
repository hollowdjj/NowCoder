package jz_II

/*
给定一个整数数组 asteroids，表示在同一行的小行星。
对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向
（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗
行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。

输入：asteroids = [5,10,-5]
输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
*/

func asteroidCollision(asteroids []int) []int {
	//用栈来模拟
	res := []int{}
	for _, a := range asteroids {
		for a < 0 && len(res) > 0 && res[len(res)-1] > 0 {
			sum := res[len(res)-1] + a
			// if sum > 0 {
			//     a = 0
			// } else if sum < 0 {
			//     res = res[:len(res)-1]
			// } else {
			//     a = 0
			//     res = res[:len(res)-1]
			// }
			if sum >= 0 {
				a = 0
			}
			if sum <= 0 {
				res = res[:len(res)-1]
			}
		}
		if a != 0 {
			res = append(res, a)
		}
	}
	return res
}
