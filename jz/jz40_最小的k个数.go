package jz

/*
给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则
最小的4个数字是1,2,3,4(任意顺序皆可)。

输入：[4,5,1,6,2,7,3,8],4
输出：[1,2,3,4]
*/

func SmallestK(input []int, k int) []int {
	//最小topk问题，用小根堆解决
	n := len(input)
	for i := n/2 - 1; i >= 0; i-- {
		adjust(input, i, n)
	}

	count := k
	for i := n - 1; i > 0; i-- {
		input[0], input[i] = input[i], input[0]
		adjust(input, 0, i)
		count--
		if count == 0 {
			break
		}
	}

	var res []int
	if k < n {
		res = input[n-k:]
	} else {
		res = input
	}

	//翻转一下
	n = len(res)
	left, right := 0, n-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

func adjust(nums []int, parent, end int) {
	child := 2*parent + 1
	for child < end {
		//在左右子节点中找到较小的那个结点
		if child+1 < end && nums[child+1] < nums[child] {
			child++
		}
		//判断是否要调整
		if nums[parent] <= nums[child] {
			break
		}
		nums[parent], nums[child] = nums[child], nums[parent]
		parent = child
		child = 2*parent + 1
	}
}
