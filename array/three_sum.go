package array

import "sort"

/*
数组中相加和为0的三元组
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c = 0？找出所有满足条件的三元组。同一个数字不能重复使用
要求：空间复杂度O(n^2), 时间复杂度O(n^2) 且三元素必须按升序排列

需要注意的是，本题只是说一个数字不能重复使用，但数组中又是可能出现重复数字的。这一点需要特别注意
*/

func ThreeSum(num []int) [][]int {
	n := len(num)
	if n < 3 {
		return [][]int{}
	}
	//将数组进行升序排序
	sort.Ints(num)
	//如果不存在负数，那么无解
	if num[0] > 0 {
		return [][]int{}
	}
	//双指针
	var res [][]int
	for i := 0; i < n-2; i++ {
		//注意，在这里，去重不是指重复子序列中只能选一个，否则我们完全可以在排序后就去除重复的元素。
		//因而，不能使用num[i+1] == num[i]来判断。原因在于，解中一个元素不能使用两次，但可以是数组
		//中的其他重复元素。例如一个数组[-40,-10,-10,0,10,20]。如果使用num[i+1] == num[i]进行完全
		//去重的话，i会直接指向第二个-10，从而漏掉[-10,-10,0]这个解。这里去重的目的在于，我们得到的
		//一个解完全有可能是[num[i],num[i+1],num[k]]。如果不跳过重复的元素的话，会造成多个重复解。
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		//固定一个位置，即i。然后使用两个指针left和right。开始时，left指向i+1，right指向
		//数组的最后一个元素。从而问题变成了判断left和right指向的值加上i指向值的是否为0。
		//此时，有三种情况：
		//1. 相加为0。此时，left和right均需要移动到非重复点
		//2. 相加和大于0。此时right--，left和i不变
		//3. 相加和小于0。此时left++，right和i不变
		left, right := i+1, len(num)-1
		for left < right {
			sum := num[left] + num[right] + num[i]
			if sum == 0 {
				//添加解
				res = append(res, []int{num[i], num[left], num[right]})
				//这里是为了去除重复解
				for left < right && num[left] == num[left+1] {
					left++
				}
				for left < right && num[right] == num[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				//注意，这里只能是right--，而不能让right一直减到不重复的元素为止。
				//原因也是，解中允许重复但在数组中不是同一个位置的元素。下面的left++
				//同理。
				right--
			} else if sum < 0 {
				left++
			}
		}
	}

	return res
}
