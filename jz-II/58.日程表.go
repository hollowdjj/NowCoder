package jz_II

import "sort"

/*
请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。

MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，
即 [start, end), 实数 x 的范围为，  start <= x < end。

当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。

每次调用MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。

请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
*/

/*
思路就是对每一个区间按照起点进行升序排序。插入的时候，使用二分法找到第一个起点大于等于end的区间。若没有找到或者找到的是
第一个区间，那么可以插入([10,20]，插入[0,10])。若找到了，但不是第一个区间，此时需要判断start是否小于等于前一个元素的end
([10,20],[20,30]，插入[25,35])，如果是则不能插入。

class MyCalendar {
public:
    MyCalendar() {

    }

    bool book(int start, int end) {
        auto iter = intervals_.lower_bound(pair{end,0});
        if(iter == intervals_.begin() || (--iter)->second <= start) {
            intervals_.insert(pair{start,end});
            return true;
        }
        return false;
    }
private:
    set<pair<int,int>> intervals_;
};
*/

type MyCalendar struct {
	books [][2]int
}

func Constructor58() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	//找到第一个起点大于等于end的区间
	index := sort.Search(len(this.books), func(i int) bool {
		return this.books[i][0] >= end
	})
	if index == 0 || this.books[index-1][1] <= start {
		this.books = append(this.books, [2]int{start, end})
		sort.Slice(this.books, func(i, j int) bool {
			if this.books[i][0] < this.books[j][0] {
				return true
			}
			return this.books[i][1] < this.books[i][1]
		})
		return true
	}
	return false
}
