package array

/*
给定一个无重叠的，按照区间起点升序排列的区间列表，在列表中插入一个新区间，如果有原区间有重合，则合并，请返回插入后的区间列表。

输入：[[2,5],[6,11]],[5,6]
输出：[[2,11]]
*/

func InsertInterval(Intervals []*Interval, newInterval *Interval) []*Interval {
	n := len(Intervals)
	if n == 0 {
		return []*Interval{newInterval}
	}

	var res []*Interval
	index := 0
	for i := 0; i < n; i++ {
		//找到newInterval的start所在区间的索引
		if newInterval.Start >= Intervals[i].Start && newInterval.Start <= Intervals[i].End {
			//随后再找newInterval的end所在区间的索引
			for j := i; j < n; j++ {
				if newInterval.End >= Intervals[j].Start && newInterval.End <= Intervals[j].End {
					//若找到了，需要判断i和j的关系
					if i == j {
						return Intervals
					}
					res = append(res, Intervals[:i]...)
					res = append(res, &Interval{Start: Intervals[i].Start, End: Intervals[j].End})
					res = append(res, Intervals[j+1:]...)
					return res
				}
			}
			//没找到newInterval的end所在区间的索引，如果新区间的end不是最大的
			res = append(res, Intervals[:i]...)
			res = append(res, &Interval{Start: Intervals[i].Start, End: newInterval.End})
			if Intervals[n-1].End > newInterval.End {
				res = append(res, Intervals[i+1:]...)
			}
			return res
		} else if Intervals[i].Start > newInterval.Start {
			index = i
			break
		}
	}

	if index > 0 {
		res = append(res, Intervals[:index]...)
		res = append(res, newInterval)
		res = append(res, Intervals[index:]...)
	} else {
		res = append(res, newInterval)
	}

	return res
}
