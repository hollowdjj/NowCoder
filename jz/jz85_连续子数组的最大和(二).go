package jz

/*
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，找到一个具有最大和的连续子数组。
1.子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组
2.如果存在多个最大和的连续子数组，那么返回其中长度最长的，该题数据保证这个最长的只存在一个
3.该题定义的子数组的最小长度为1，不存在为空的子数组，即不存在[]是某个数组的子数组
4.返回的数组不计入空间复杂度计算

输入：[1,-2,3,10,-4,7,2,-5]
输出：[3,10,-4,7,2]
说明：经分析可知，输入数组的子数组[3,10,-4,7,2]可以求得最大和为18，故返回[3,10,-4,7,2]
*/

func FindGreatestSumOfSubArray2(array []int) []int {
	//dp[i]表示以array[i]结尾的连续子数组的最大和，状态转移方程为：
	//dp[i] = max(dp[i-1]+array[i],array[i])，可进行状态压缩。
	prev, curr := array[0], 0
	n := len(array)
	resLeft, resRight := 0, 0
	currLeft, currRight := 0, 0
	maxSum := prev
	for i := 1; i < n; i++ {
		if prev+array[i] >= array[i] {
			curr = prev + array[i]
			currRight = i
		} else {
			curr = array[i]
			currLeft, currRight = i, i
		}
		if curr > maxSum || curr == maxSum && currRight-currLeft+1 > resRight-resLeft+1 {
			//当有更大和，或者最大和一样但是区间更长时，更新区间
			resLeft, resRight = currLeft, currRight
			maxSum = curr
		}
		prev = curr
	}
	return array[resLeft : resRight+1]
}
