package jz

/*
一个整型数组里除了一个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字
*/

func FindNumAppearOnce1(nums []int) int {
	//利用异或的性质
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

/*
一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字
*/

func FindNumAppearOnce2(nums []int) []int {
	//需要注意的是，这道题让我找出两个只出现了一次的数字。若还是用上一题的做法，最后得到的结果是
	//a⊕b，而我们需要分别得到a和b。一个可能的方法是对数组中的数字进行分组，使得：
	// 					a⊕c⊕d⊕c⊕d.... = a   且 	b⊕x⊕y⊕x⊕y... = b
	//因此，现在的问题就是如何进行分组。a和b的二进制值是不同的，因此我们只需要找到这两个数从低位
	//到高位中的第一个不相同的位，然后将这个位置为1，其余位为0的值就能够区分出a,b两个数了。我们可
	//以根据异或的结果进行移位，得到这个数。
	temp := 0
	for _, v := range nums {
		temp ^= v
	}
	mask := 1
	for temp&mask == 0 {
		mask <<= 1
	}

	a, b := 0, 0
	for _, v := range nums {
		if mask&v == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}

	if a < b {
		return []int{a, b}
	}
	return []int{b, a}
}

/*
一个整型数组里除了一个数字只出现一次，其他的数字都出现了k次。请写程序找出这两个只出现一次的数字
*/
func FindNumAppearOnce3(nums []int32, k int) int32 {
	//由于其他数字都出现了k次，因为不能再适用于异或求解了。首先定义一个数组，用于存放数组中各个
	//数字中的每一位按位相加的结果。从而可以通过哪些不是k的倍数的位，反推回只出现了一次的数字。
	temp := make([]int32, 32)
	for _, v := range nums {
		for i := 0; i < 32; i++ {
			temp[i] += (v >> i) & 1
		}
	}

	var res int32
	for i := 0; i < 32; i++ {
		if temp[i]%int32(k) == 1 {
			res |= 1 << i
		}
	}

	return res
}
