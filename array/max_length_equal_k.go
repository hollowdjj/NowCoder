package array

/*
和为k的连续子数组

给定一个无序数组 arr , 其中元素可正、可负、可0。给定一个整数 k ，求 arr 所有连续子数组
中累加和为k的最长连续子数组长度。保证至少存在一个合法的连续子数组。
[1,2,3]的连续子数组有[1,2]，[2,3]，[1,2,3] ，但是[1,3]不是

输入：[1,-2,1,1,1],0
输出：3
*/

func MaxLengthEqualK(arr []int, k int) int {
	//使用一个哈希表保存a[0...i]子数组的和。其中key为前缀和，value为该前缀和第一次出现时的位置i，这样才能保证结果是最长的。
	dic := make(map[int]int)
	n := len(arr)
	sum := 0
	res := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if _, ok := dic[sum]; !ok {
			dic[sum] = i
		}
		if j, ok := dic[sum-k]; ok {
			//设sumi为arr[0...i]的和,sumj为arr[0...j]的和。那么arr[j+1....i]的和为sumi-sumj
			//所以，当sumi-sumj == k时，我们就可以更新最大长度res=max(res,i-j)
			res = max(res, i-j)
		}
	}

	return res
}
