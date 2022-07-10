package jz_II

import "sort"

/*
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差
并以分钟数表示。

输入：timePoints = ["23:59","00:00"]
输出：1
*/

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	//超过1440个，那么必有重复时间点，直接返回0
	if n > 1440 {
		return 0
	}
	//以0点为准得到每一个时间距离其的分钟数并排序
	temp := make([]int, 0, n)
	for _, t := range timePoints {
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minutes := hour*60 + int(t[3]-'0')*10 + int(t[4]-'0')
		if hour == 0 && minutes == 0 {
			minutes = 1440
		}
		temp = append(temp, 1440-minutes)
	}
	sort.Ints(temp)
	//这里要将最小值加上1440然后添加到数组尾部。用来应对最小时间差
	//是0点左右两边的时间点的情况，例如["05:31","22:08","00:35"]
	temp = append(temp, temp[0]+1440)
	res := 1440
	for i := 0; i < len(temp)-1; i++ {
		if t := temp[i+1] - temp[i]; t < res {
			res = t
		}
	}
	return res
}
