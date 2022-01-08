package array

/*
牛妹给了牛牛一个长度为n的下标从0开始的正整型数组a ，粗心的牛牛不小心把其中的一些数字删除了。
假如a_i被删除了,则a_i=0。对于所有被删除的数字，牛牛必须选择一个正整数填充上。现在牛牛想知道有多少种填充方案使得：
a0≤a1≤a2...≤an-1，且1≤ai≤k
函数传入一个下标从0开始的数组a和一个正整数k ，请返回合法的填充方案数对 10^9+7取模的值,保证不存在方案数为0的数据。

输入：[0,4,5],6
返回值：4
说明：所有的合法填充方案是：[1,4,5],[2,4,5],[3,4,5],[4,4,5]，共4种。
*/

func FillArray(a []int, k int) int {
	//以数组2,0,0,0,5为例。易知，对每一个0序列而言，填充的方案数仅仅取决于0序列的长度
	//以及0序列的上下界。因此，可以使用动态规划。dp[i][j]，其中i为0序列可填充数的数量，为
	//ub-lb+1；j为0序列的长度减一。
	//0序列的长度m=3，上界为5，下届为2，因此方案数为4^2 + 3^2 + 2^2 + 1^2
	//通过数学归纳法，可得

	length := len(a)
	res := 1

	//left为0序列第一个数的位置，right为0序列最后一个数的后一个位置
	left := 0
	for left < length {
		if a[left] != 0 {
			left++
			continue
		}
		//往后找到0序列的最后一个元素的位置right并得到合法填充值的上下限
		//这里存在几种情况：1. right <= length-1 2.right == length
		right := left + 1
		for ; right < length; right++ {
			if a[right] != 0 {
				break
			}
		}
		//得到可填充数字的上下限
		lb, ub := 1, k
		m := right - left //一共有m个0元素要填充的个数
		if left >= 1 {
			lb = a[left-1]
		}
		if right < length {
			ub = a[right]
		} else {
			m = right - left + 1
		}
		//当前0元素序列的填充方案
		n := ub - lb + 1
		temp := 0
		for i := 0; i < n; i++ {
			temp += pow(i, m-1)
		}
		res *= temp
		//res %= 1000000007

		left = right + 1
	}

	return res % 1000000007
}

func pow(x int, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}
