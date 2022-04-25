package jz

/*
把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑
数。求按从小到大的顺序的第 n个丑数。
*/

func GetUglyNumber(index int) int {
	if index < 6 {
		return index
	}
	//按照定义一个丑数=2^x * 3^y * 5^z。因此，我们需要定义一个数组用于存放
	//按照大小顺序排列的丑数。然后让最后一个元素分乘以2、3、5，然后取小的那个
	//然而，需要注意的是，后一轮的最小数不一定比前一轮较大的两个数小，说以需要有一种方式
	//保存前一轮，甚至前n轮没有被选上的数。为了实现这一目的，使用3个指针i2,i3,i5，分别指向
	//前面n轮中乘2后，乘3后以及乘5后没被选上的数。
	nums := make([]int, index)
	nums[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < index; i++ {
		nums[i] = min(min(nums[i2]*2, nums[i3]*3), nums[i5]*5)
		//注意，这里不能写if -else ，因为可能存在nums[i2]*2 == nums[i3]*3这种情况
		//那么在下一次循环中，就会造成v中有重复元素出现
		if nums[i] == nums[i2]*2 {
			i2++
		}
		if nums[i] == nums[i3]*3 {
			i3++
		}
		if nums[i] == nums[i5]*5 {
			i5++
		}
	}
	return nums[index-1]
}
