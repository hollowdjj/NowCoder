package codetop

/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	//arr = [1,2,3,4]的差分数组为num = [1,1,1,1]。即num[i] = arr[i] - arr[i-1]
	//差分数组的前缀和就是原数组。当我们对原数组区间[l,r]中的元素都加x后，为了维护差分数组
	//的正确性，num[l] += x，num[r+1] -= x。这是因为，num[l] += x，相当于把arr[l:]中的
	//所有数都加x，所以num[r+1]要减x，这样arr[r+1]与arr[r]的差就少了x，差分数组才是正确的。
	nums := make([]int, n)
	//构建差分数组
	for _, book := range bookings {
		nums[book[0]-1] += book[2]
		if book[1] < n {
			nums[book[1]] -= book[2]
		}
	}
	//计算差分数组的前缀和
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
