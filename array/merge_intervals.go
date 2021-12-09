package array

/*
合并区间
给出一组区间，请合并所有重叠的区间。
要求：空间复杂度O(n)，时间复杂度O(nlogn)
进阶：空间复杂度O()

示例：
输入：[[10,30],[20,60],[80,100],[150,180]]
输出：[[10,60],[80,100],[150,180]]
*/

//MergeIntervals 归并排序解决合并区间问题
func MergeIntervals(intervals []*Interval) []*Interval {
	if len(intervals) < 2 {
		return intervals
	}
	//按照左区间的大小进行升序排列
	ordered := reorderIntervals(intervals)
	//遍历并合并区间
	var res []*Interval
	res = append(res, ordered[0])
	j := 0
	for i := 1; i < len(ordered); i++ {
		if ordered[i].Start <= res[j].End && ordered[i].End > res[j].End {
			//合并重复区间
			res[j].End = ordered[i].End
		} else if ordered[i].Start > res[j].End {
			//区间不相交，则直接添加
			res = append(res, ordered[i])
			j++
		}
		//剩下的这种情况是ordered[i]在res[j]区间里面。这种情况不需要任何操作
	}
	return res
}

func reorderIntervals(intervals []*Interval) []*Interval {
	//递归终止条件
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	mid := n / 2
	return merge(reorderIntervals(intervals[:mid]), reorderIntervals(intervals[mid:]))
}

//合并两个升序区间数组
func merge(arr1, arr2 []*Interval) []*Interval {
	var res []*Interval
	n1, n2, i, j := len(arr1), len(arr2), 0, 0

	for i < n1 && j < n2 {
		if arr1[i].Start < arr2[j].Start {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	if i == n1 {
		res = append(res, arr2[j:]...)
	} else if j == n2 {
		res = append(res, arr1[i:]...)
	}

	return res
}
