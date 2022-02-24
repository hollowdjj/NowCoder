package array

/*
最大差值
有一个长为 n 的数组 A ，求满足 0 ≤ a ≤ b < n 的 A[b] - A[a] 的最大值。
给定数组 A 及它的大小 n ，请返回最大差值。
*/

func MaxDifference(A []int, n int) int {
	res := 0
	minVal := A[0]
	for i := 1; i < n; i++ {
		minVal = min(minVal, A[i]) //A[0..i]中的最小值
		res = max(res, A[i]-minVal)
	}
	return res
}
