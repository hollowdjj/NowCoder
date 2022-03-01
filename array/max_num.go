package array

import (
	"sort"
	"strconv"
	"strings"
)

/*
最大数

给定一个长度为n的数组nums，数组由一些非负整数组成，现需
要将他们进行排列并拼接，每个数不可拆分，使得最后的结果最大，返回值需要是string类型，否则可能会溢出。

输入：[2,20,23,4,8]
输出："8423220"
进阶：时间复杂度O(nlogn)，空间复杂度：O(n)
*/

func MaxNum(nums []int) string {
	//本质上就是排序+自定义两数的比较函数。以2和20两个数为例，"220" > "202"的，故如果是将
	//数组按照降序排序的话，2应该在20前面。

	//先将数组中的每个数转换成字符串
	n := len(nums)
	numStr := make([]string, 0, n)
	for _, num := range nums {
		numStr = append(numStr, strconv.Itoa(num))
	}

	//排序
	sort.Slice(numStr, func(i, j int) bool {
		if numStr[i]+numStr[j] > numStr[j]+numStr[i] {
			//numStr[i]在前面
			return true
		}
		return false
	})

	res := strings.Join(numStr, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
