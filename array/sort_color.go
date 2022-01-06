package array

/*
给定一个包含红色，白色，蓝色，一同 n 个元素的数组，对其进行排序使得相同的颜色相邻并且按照 红色，白色，蓝色的顺序排序。
数组中 0 代表红色，1 代表白色，2 代表蓝色。

输入：[0,2,1]
输出：[0,1,2]
*/

func SortColor(colors []int) []int {
	//利用快排的思想，将大于1的放在右边，小于1的放在左边。具体的做法为：
	//首先，定义less,more两个指针。其中less指向小于区域的右边界，而more
	//指向大于区域的左边界。这两个指针分别初始化为-1和len(colors)。
	//随后，从curr=0开始：
	//1. 如果colors[curr] == 0，那么将swap[colors[curr],colors[less+1]]，然后less++,curr++
	//2. 如果colors[curr] == 1，那么curr++
	//3. 如果colors[curr] == 2，那么swap[colors[curr],colors[more-1]]，然后more--。需要注意的是，
	//此时curr不能++，因为colors[more-1]可能等于0，此时还需要继续往前交换
	less, more, curr := -1, len(colors), 0

	for curr < more {
		switch v := colors[curr]; v {
		case 2:
			more--
			colors[curr], colors[more] = colors[more], colors[curr]
		case 0:
			less++
			colors[curr], colors[less] = colors[less], colors[curr]
			fallthrough
		default:
			curr++
		}
	}

	return colors
}
