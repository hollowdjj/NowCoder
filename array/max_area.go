package array

/*
给定一个数组height，长度为n，每个数代表坐标轴中的一个点的高度，height[i]是在第i点的高度，请问，从中选2个高度与x轴组成的容器最多能容纳多少水
1.你不能倾斜容器
2.当n小于2时，视为不能形成容器，请返回0
3.数据保证能容纳最多的水不会超过整形范围
*/

func MaxArea(height []int) int {
	//双指针，左指针最开始指向数组头，右指针最开始指向数组尾。容器能盛水的面积等于最短边的边长乘以
	//容器宽度。每次，向内移动一条边，容器宽度一定是减小的。若将短的那条边向内移动，面积有增大的可能
	//但，将长的那条边想内移动，则面积绝无增大的可能
	left, right := 0, len(height)-1

	res := 0
	for left < right {
		res = MaxOfTwoInt(res, MinOfTwoInt(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}
