package jz

/*
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；
针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}，
{2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
*/

func MaxInWindows(num []int, size int) []int {
	n := len(num)
	//一个降序数组，第一个元素为num[0..i]的最大值
	deque := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(deque) > 0 && num[deque[len(deque)-1]] < num[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		//判断首元素是否还在滑动窗口内 0 1 2
		if deque[0]+size <= i {
			deque = deque[1:]
		}
		//滑动窗口刚开始时必须累计到size个
		if i+1 >= size {
			res = append(res, num[deque[0]])
		}
	}
	return res
}
