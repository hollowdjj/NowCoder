package codetop

/*
给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
由于答案可能很大，因此 返回答案模 10^9 + 7 。
输入：arr = [3,1,2,4]
输出：17
解释：
子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
*/

func sumSubarrayMins(arr []int) int {
	//对元素E=arr[i]而言，若E是arr[left:right]中的最小值，那么arr[left:right]的
	//所有连续子数组的最小值均为E。从而现在有两个问题需要解决：
	//1. arr[left:right]能产生多少个包含E的连续子数组。我们可以通过枚举左右边界的形式来看。
	//连续子数组的左边界必须在[left..i]中选取，一共(i-left+1)个；右边界必须在[i..right中]
	//选，一共有(right-i+1)个，因此一共可以产生(i-left+1)*(right-i+1)个包含E的连续子数组。
	//2. arr[left:right]的left和right如何确定。其实就是要找到arr[i]右边和左边第一个小于它的数
	//那么可以使用单调栈实现
	var s []int
	i, n := 0, len(arr)
	var res int64 = 0
	for i < n {
		if len(s) == 0 || arr[i] > arr[s[len(s)-1]] {
			s = append(s, i)
			i++
		} else {
			//对于E=arr[s[len(s)-1]]而言，arr[i] < E && arr[s[len(s)-2]] < E
			//left = s[len(s)-2]] + 1,right = i-1
			top := s[len(s)-1]
			s = s[:len(s)-1]
			E := arr[top]
			left, right := 0, i-1
			if len(s) > 0 {
				left = s[len(s)-1] + 1
			}
			res += (int64(E) * int64(top-left+1) * int64(right-top+1)) % (1e9 + 7)
		}
	}
	for len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		E := arr[top]
		left, right := 0, n-1
		if len(s) > 0 {
			left = s[len(s)-1] + 1
		}
		res += (int64(E) * int64(top-left+1) * int64(right-top+1)) % (1e9 + 7)
	}
	return int(res % (1e9 + 7))
}
