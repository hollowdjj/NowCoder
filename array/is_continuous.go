package array

import "sort"

/*
扑克牌顺子
现在有2副扑克牌，从扑克牌中随机五张扑克牌，我们需要来判断一下是不是顺子。
有如下规则：
1. A为1，J为11，Q为12，K为13，A不能视为14
2. 大、小王为 0，0可以看作任意牌
3. 如果给出的五张牌能组成顺子（即这五张牌是连续的）就输出true，否则就输出false。
4.数据保证每组5个数字，每组最多含有4个零，数组的数取值为 [0, 13]

要求：空间复杂度 O(1)O(1)，时间复杂度 O(nlogn)O(nlogn)，本题也有时间复杂度 O(n)O(n) 的解法

输入：[6,0,2,0,4]
返回值：true
说明：中间的两个0一个看作3，一个看作5 。即：[6,3,2,5,4] 这样这五张牌在[2,6]区间连续，输出true
*/

func IsContinuous(numbers []int) bool {
	//排序
	sort.Ints(numbers)
	//遍历排序后的数组，统计0的个数，以及相邻两数的差
	n := len(numbers)
	i := 0
	for numbers[i] == 0 {
		i++
	}
	zeroCount := i
	need := 0
	for ; i < n-1; i++ {
		if numbers[i] == numbers[i+1] {
			return false
		}
		need += numbers[i+1] - numbers[i] - 1
		if need > zeroCount {
			return false
		}
	}

	return true
}
