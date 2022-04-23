package jz

/*
给定一个数组 A[0,1,...,n-1] ,请构建一个数组 B[0,1,...,n-1] ,其中 B 的元素 B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]
（除 A[i] 以外的全部元素的的乘积）。程序中不能使用除法。
（注意：规定 B[0] = A[1] * A[2] * ... * A[n-1]，B[n-1] = A[0] * A[1] * ... * A[n-2]）
对于 A 长度为 1 的情况，B 无意义，故而无法构建，用例中不包括这种情况。

输入：[1,2,3,4,5]
输出：[120,60,40,30,24]
*/

func multiply(A []int) []int {
	//令left[i]表示A[0]*A[1]*....*A[i-1]，而right[i] = A[i+1]*....*A[n-1]
	//从而可得B[i] = left[i]*right[i]。其中：
	//left[i] = left[i-1]*A[i-1]; right[i] = right[i+1]*A[i+1]
	n := len(A)
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * A[i-1]
	}

	res := make([]int, n)
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = left[i] * right
		right *= A[i]
	}
	return res
}
