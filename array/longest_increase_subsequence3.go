package array

/*
最长上升子序列(三)
给定数组 arr ，设长度为 n ，输出 arr 的最长上升子序列。
（如果有多个答案，请输出其中 按数值(注：区别于按单个字符的ASCII码值)进行比较的 字典序最小的那个）

输入：[1,2,8,6,4]
输出：[1,2,4]
说明：其最长递增子序列有3个，（1，2，8）、（1，2，6）、（1，2，4）其中第三个
按数值进行比较的字典序 最小，故答案为（1，2，4）
*/

func LIS3(arr []int) []int {
	//dp[i]表示以arr[i]结尾的最长上升子序列。故，状态转移方程为：
	//dp[i] = Max{dp[j] where arr[i]>arr[j]} + 1，其中j=0....i-1
	n := len(arr)
	if n < 2 {
		return arr
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	maxLength := 0
	for i := 1; i < n; i++ {
		//状态转移的时间复杂度为O(n)，因此总的时间复杂度为O(n^2)，会超时
		temp := 0
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				temp = max(temp, dp[j])
			}
		}
		dp[i] = temp + 1
		maxLength = max(maxLength, dp[i])
	}

	res := make([]int, maxLength)
	j := maxLength
	//倒叙遍历才能保证长度相同的情况下是字典序最小的
	for i := n - 1; j > 0; i-- {
		if dp[i] == j {
			j--
			res[j] = arr[i]
		}
	}
	return res
}

func LIS3Advance(arr []int) []int {
	//我们需要优化一下动态规划中状态转移时的复杂度。具体的做法是维护一个辅助数组end。
	//其中end[k]表示所有上升子序列长度为k+1的子序列最后一个元素的最小值。例如：有上
	//升子序列[1,2,3] [0,1,2] [1,2,4]，那么end[2]=2。end数组的维护方法如下：
	//1. 当arr[i] > end[len(end)-1]时，说明end[len(end)-1]与arr[i]又构成了一个上升
	//序列，故dp[i] = len(end)+1。append(end,arr[i])
	//2. 当arr[i] <= end[len(end)-1]时，我们要在end数组中找到第一个大于等于arr[i]的
	//数的位置k，并让end[k]=arr[i]

	//dp[i]表示以arr[i]结尾的子串的最长上升子序列长度，那么状态转移方程为：
	//dp[i] = Max{dp[j]}+1 其中j<i
	n := len(arr)
	if n < 2 {
		return arr
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	end := make([]int, 1)
	end[0] = arr[0]
	for i := 1; i < n; i++ {
		//遍历dp[0:i]找到结尾小于arr[i]并且长度最大的递增序列。此时状态转移的
		//时间复杂度为O(n)，总体的时间复杂度为O(n^2)会超时。因此，需要优化状态
		//转移部分。维护一个辅助数组end。end[k]为所有长度为k+1的子序列的尾元素中的最
		//小值。故当前最长递增序列的长度就是len(end)，且end是一个递增数组维护的步骤为：
		//当arr[i] > end[len(end)-1]时，说明end[len(end)-1]与arr[i]又构成了一个上升
		//序列，故dp[i] = len(end)+1。append(end,arr[i])
		if arr[i] > end[len(end)-1] {
			dp[i] = len(end) + 1
			end = append(end, arr[i])
		} else {
			//此时我们需要找到end数组中第一个大于等于arr[i]的元素的位置k，并将其修改为arr[i]
			//表示我们将长度为k+1的递增子序列的最后一个元素修改为了一个更小的值。这是一个贪心
			//算法。因此，要想使得上升序列尽量长，那么就要使得序列上升的足够慢，就是说两两元素
			//之间的差值要尽量小。由于end是一个递增数组，故使用二分法即可。
			left, right := 0, len(end)
			for left < right {
				mid := (right + left) / 2
				if end[mid] < arr[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			end[left] = arr[i]
			dp[i] = left + 1
		}
	}

	//倒序遍历end
	maxLength := len(end)
	res := make([]int, maxLength)
	for i := n - 1; i >= 0; i-- {
		if dp[i] == maxLength {
			res[maxLength-1] = arr[i]
			maxLength--
		}
	}
	return res
}
