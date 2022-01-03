package array

import "math"

/*
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，找到一个具有最大和的连续子数组。
1.子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组
2.如果存在多个最大和的连续子数组，那么返回其中长度最长的，该题数据保证这个最长的只存在一个
3.该题定义的子数组的最小长度为1，不存在为空的子数组，即不存在[]是某个数组的子数组
4.返回的数组不计入空间复杂度计算

要求:时间复杂度O(n)，空间复杂度O(n)
进阶:时间复杂度O(n)，空间复杂度O(1)
*/

func FindGreatestSumOfSubArray(array []int) []int {
	n := len(array)
	if n == 1 {
		return array
	}

	resLeft, resRight := 0, 0
	currLeft, currRight := 0, 0
	res := math.MinInt //连续子数组最大和
	sum := 0           //currLeft到currRight间所有元素的和

	for currRight < n {
		sum += array[currRight]
		if sum == res {
			//如果等于之前的解，那么比较长度
			if currRight-currLeft >= resRight-resLeft {
				resLeft, resRight = currLeft, currRight
			}
			res = sum
		} else if sum > res {
			//如果大于之前的解，直接更新最终解
			resLeft, resRight = currLeft, currRight
			res = sum
		}

		//一旦前i个元素的和小于0，那么就要移动currLeft，因为此时后面只要有一个正数，那么
		//就存在新的解。如果后面数全是负数，那么当前解就是最终的解了
		if sum < 0 {
			currLeft = currRight + 1
			sum = 0
		}
		currRight++
	}

	return array[resLeft : resRight+1]
}
