package easy

import "fmt"

/*
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
数据范围: 1 <= n <= 10^5
要求:时间复杂度为O(n)，空间复杂度为 O(1)
*/

func FindGreatestSumOfSubArray(array []int) int {
	n := len(array)
	//dp数组的定义为: dp[i]表示以array[i]元素结尾的连续子数组的最大和，最后返回的结果为dp[n-1]
	dp := make([]int, n)
	dp[0] = array[0]
	//选择
	for i := 1; i < n; i++ {
		dp[i] = maxOfTwoInt(array[i], dp[i-1]+array[i])
	}
	//遍历dp table得到最大值
	return findMaxValInSlice(dp)
}

func maxOfTwoInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func findMaxValInSlice(slice []int) int {
	res := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > res {
			res = slice[i]
		}
	}
	return res
}

func TestSubArrayMaxSum() {
	fmt.Println(FindGreatestSumOfSubArray([]int{1, -2, 3, 10, -4, 7, 2, -5}))
}
