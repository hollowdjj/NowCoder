package jz

import "sort"

/*
现在有2副扑克牌，从扑克牌中随机五张扑克牌，我们需要来判断一下是不是顺子。
有如下规则：
1. A为1，J为11，Q为12，K为13，A不能视为14
2. 大、小王为 0，0可以看作任意牌
3. 如果给出的五张牌能组成顺子（即这五张牌是连续的）就输出true，否则就输出false。
4.数据保证每组5个数字，每组最多含有4个零，数组的数取值为 [0, 13]

输入：[6,0,2,0,4]
输出：true
*/

func IsContinuous(numbers []int) bool {
	sort.Ints(numbers)
	//找到第一个不为0的元素的索引
	i := 0
	for ; i < len(numbers); i++ {
		if numbers[i] != 0 {
			break
		}
	}
	//判断有无重复非零元素
	for j := i; j < len(numbers)-1; j++ {
		if numbers[j] == numbers[j+1] {
			return false
		}
	}

	//非零区间首元素和尾元素的差必须小于5，否则无法构成顺子，因为一共就5张牌：
	//[0,0,2,4,6] true [0,0,2,4,7] false
	return numbers[4]-numbers[i] < 5
}
