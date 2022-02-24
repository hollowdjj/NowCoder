package others

/*
跳跃游戏(一)
给定一个非负整数数组nums，假定最开始处于下标为0的位置，数组里面的每个元
素代表下一跳能够跳跃的最大长度。如果能够跳到数组最后一个位置，则返回true，否则返回false。

输入：[2,1,3,3,0,0,100]
返回：true
说明：首先位于nums[0]=2，然后可以跳2步，到nums[2]=3的位置，再跳到nums[3]=3的位
置，再直接跳到nums[6]=100，可以跳到最后，返回true
*/

func JumpGame(nums []int) bool {
	pos := 0 //当前能跳到的最远位置
	n := len(nums)
	for i := 0; i < n; i++ {
		if pos < i {
			//如果当前能跳到的最远的地方都到不了i的话，说明不可能跳到最后一个位置
			return false
		}
		if pos >= n-1 {
			return true
		}
		pos = max(pos, i+nums[i])
	}
	return true
}
