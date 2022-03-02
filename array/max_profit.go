package array

/*
假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
2.如果不能获取到任何利润，请返回0
3.假设买入卖出均无手续费

输入：[8,9,2,5,4,7,1]
输出：5
说明：在第3天(股票价格 = 2)的时候买入，在第6天(股票价格 = 7)的时候卖出，最大利润 = 7-2 = 5 ，
      不能选择在第2天买入，第3天卖出，这样就亏损7了；同时，你也不能在买入前卖出股票。

要求：空间复杂度 O(1)，时间复杂度 O(n)
*/

func MaxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}

	minVal, maxVal := prices[0], -1
	for i := 1; i < n; i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		}
		maxVal = max(maxVal, prices[i]-minVal)
	}
	return maxVal
}
