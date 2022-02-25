package string

import "strconv"

/*
字符串出现次数的TopK问题
给定一个字符串数组，再给定整数 k ，请返回出现次数前k名的字符串和对应的次数。
返回的答案应该按字符串出现频率由高到低排序。如果不同的字符串有相同出现频率，按字典序排序。
对于两个字符串，大小关系取决于两个字符串从左到右第一个不同字符的 ASCII 值的大小关系。
比如"ah1x"小于"ahb"，"231"<”32“
字符仅包含数字和字母

输入：["123","123","231","32"],2
输出：[["123","2"],["231","1"]]
说明： "123"出现了2次，记["123","2"]，"231"与"32"各出现1次，但
是"231"字典序在"32"前面，记["231","1"]，最后返回[["123","2"],["231","1"]]
*/

type pair struct {
	str   string
	count int
}

func (p pair) biggerThan(rhs pair) bool {
	if p.count > rhs.count {
		return true
	}

	//次数相同的时候比较字典序
	return p.count == rhs.count && p.str < rhs.str
}

func TopKStrings(strings []string, k int) [][]string {
	//先使用一个哈希表记录每个字符串出现的次数
	dic := make(map[string]int)
	for _, s := range strings {
		dic[s]++
	}

	//将哈希表中的元素保存到一个数组中以进行堆排序
	n := len(dic)
	arr := make([]pair, 0, n)
	for s, v := range dic {
		arr = append(arr, pair{s, v})
	}

	//构建一个大顶堆。即从最后一个非叶子结点开始调整
	for i := n/2 - 1; i >= 0; i-- {
		adjust(arr, i, n)
	}

	//堆排序好最大的k个元素后就返回
	res := make([][]string, 0, k)
	for i := n - 1; i > n-1-k; i-- {
		//交换arr[i]和arr[0]
		arr[i], arr[0] = arr[0], arr[i]
		res = append(res, []string{arr[i].str, strconv.Itoa(arr[i].count)})
		//调整arr[0:i]
		adjust(arr, 0, i)
	}
	return res
}

func adjust(arr []pair, parentIndex int, end int) {
	childIndex := 2*parentIndex + 1
	for childIndex < end {
		//得到左右孩子中的较大值
		if childIndex+1 < end && arr[childIndex+1].biggerThan(arr[childIndex]) {
			childIndex++
		}
		//如果父亲大于最大的孩子，那么不需要调整
		if arr[parentIndex].biggerThan(arr[childIndex]) {
			break
		}
		//否则交换父亲和孩子，并递归调整
		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
}
