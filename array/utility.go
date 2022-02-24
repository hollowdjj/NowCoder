package array

import (
	"fmt"
	"strings"
)

//Interval 区间
type Interval struct {
	Start int
	End   int
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%d,%d]", i.Start, i.End)
}

func (i *Interval) Equal(another *Interval) bool {
	if i.Start == another.Start && i.End == another.End {
		return true
	}

	return false
}

//PrintIntervals 打印区间数组
func PrintIntervals(intervals []*Interval) string {
	n := len(intervals)
	if n < 1 {
		return "[]"
	}
	temp := make([]string, n)
	for i, v := range intervals {
		temp[i] = v.String()
	}
	res := strings.Join(temp, ",")

	return "[" + res + "]"
}

//SliceToInterval 将一个二维数组转成区间数组
func SliceToInterval(slice [][2]int) []*Interval {
	n := len(slice)
	if n < 1 {
		return nil
	}
	res := make([]*Interval, n)
	for i := 0; i < n; i++ {
		res[i] = &Interval{Start: slice[i][0], End: slice[i][1]}
	}

	return res
}

//EqualIntervalSlice 判断两区间数组是否相等
func EqualIntervalSlice(arr1, arr2 []*Interval) bool {
	n1, n2 := len(arr1), len(arr2)
	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		if !arr1[i].Equal(arr2[i]) {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
