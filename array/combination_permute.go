package array

import (
	"sort"
)

/*
排列
*/

//不含重复元素的全排列，返回其所有可能的全排列
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func Permute1(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	used := make([]bool, n)
	var backtrack func(curr []int)
	backtrack = func(curr []int) {
		//递归终止条件
		if len(curr) == n {
			temp := make([]int, n)
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			//全排列的话需要使用一个used数组记录已经被选择了的数
			if used[i] {
				continue
			}
			used[i] = true
			curr = append(curr, nums[i])
			backtrack(curr)
			//回溯
			used[i] = false
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{})
	return res
}

//包含重复元素的全排列，返回其所有可能的排列
//输入：nums = [1,1,2]
//输出：[1,1,2],[1,2,1],[2,1,1]
func PermuteUni(nums []int) [][]int {
	n := len(nums)
	//必须先排序，使得重复的元素相邻
	sort.Ints(nums)
	res := make([][]int, 0)
	used := make([]bool, n)
	var backtrack func(curr []int)
	backtrack = func(curr []int) {
		//递归终止条件
		if len(curr) == n {
			temp := make([]int, n)
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			//这里的!used[i-1]是有讲究的。以[1,1,2]为例，递归数如下图所示：
			//								开始
			//                 1              1              2
			//        [1,1]       [1,2]
			//       [1,1,2]     [1,2,1]
			//在第一层中，取第一个1，递归完成后，得到解[1,1,2]和[1,2,1]。回溯到第一层后，我们不能选择第二个1
			//否则就会重新得到[1,1,2]和[1,2,1]这两个解。因此，必须要屏蔽掉这种情况。这也是开始时要把数组排一
			//下序的原因。当回溯到第一层的时候used数组为[0,0,0]，所以就屏蔽掉了这种情况。也就是说，这里是一个同层
			//剪枝。需要注意的是，这里换成used[i-1]也是可以的，只不过是树枝剪枝，效率低于前者。
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			curr = append(curr, nums[i])
			backtrack(curr)
			//回溯
			used[i] = false
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{})
	return res
}

/*
组合
*/

//不带重复数字的数组[1,2,3,4,5,...n]中所有k个数的组合
//输入：n = 4 k = 2
//输出：[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
func Combination(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrack func(i int, curr []int)
	backtrack = func(i int, curr []int) {
		if len(curr) == k {
			temp := make([]int, k)
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for j := i; j <= n; j++ {
			curr = append(curr, j)
			backtrack(j+1, curr)
			//回溯
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(1, []int{})
	return res
}

//组合总和(一)
//给你一个无重复元素的整数数组candidates和一个目标整数target，找出candidates中可以使数字和为目标数
//target的所有不同组合，并以列表形式返回。你可以按任意顺序返回这些组合。
//candidates中的同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。
func CombinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	res := make([][]int, 0)
	var backtrack func(curr []int, i int, rest int)
	backtrack = func(curr []int, i int, rest int) {
		//递归终止条件
		if rest == 0 {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		if rest < 0 {
			return
		}

		for j := i; j < n; j++ {
			curr = append(curr, candidates[j])
			//同一个数字可以重复使用，所以这里传参还是j
			backtrack(curr, j, rest-candidates[j])
			//回溯
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{}, 0, target)
	return res
}

//组合总和(二)
//给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合
//candidates中每个数字在每个组合中只能使用一次。candidates中可能有重复元素
//输入: candidates = [2,5,2,1,2], target = 5
//输出: [1,2,2], [5]
func CombinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	res := make([][]int, 0)
	//排序，使得相同的元素相邻
	sort.Ints(candidates)
	var backtrack func(i int, curr []int, rest int)
	backtrack = func(i int, curr []int, rest int) {
		if rest == 0 {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		if rest < 0 {
			return
		}

		for j := i; j < n; j++ {
			//输入：[10,1,2,7,6,1,5],8时，排序后为[1,1,2,5,6,7,10]。避免出现两次[1,2,5]
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			curr = append(curr, candidates[j])
			backtrack(j+1, curr, rest-candidates[j])
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0, []int{}, target)
	return res
}

//组合总和(三)
//找出所有相加之和为n的k个数的组合，且满足下列条件：只使用数字1到9且每个数字最多使用一次
//返回所有可能的有效组合的列表。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//输入：k = 3,n = 7
//输出：[[1,2,4]]
//解释：1+2+4 = 7
func CombinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtrack func(curr []int, i, rest int)
	backtrack = func(curr []int, i, rest int) {
		//递归终止条件
		length := len(curr)
		if length == k && rest == 0 {
			temp := make([]int, length)
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		if length > k || rest < 0 {
			return
		}

		for j := i; j <= 9; j++ {
			curr = append(curr, j)
			backtrack(curr, j+1, rest-j)
			//回溯
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{}, 1, n)
	return res
}

//组合总和(四)
//给你一个由不同整数组成的数组nums，和一个目标整数target。请你从nums中找出并返回总和为target的元素组合的个数。
//输入：nums = [1,2,3], target = 4
//输出：7
//解释：
//所有可能的组合为： (1, 1, 1, 1) (1, 1, 2) (1, 2, 1) (1, 3) (2, 1, 1) (2, 2) (3, 1)
//请注意，顺序不同的序列被视作不同的组合。
func CombinationSum4(nums []int, target int) int {
	//这道题回溯会超时，本质上就是找零钱，利用动态规划求解。dp[i]表示凑出i的组合数。
	//状态转移方程为：dp[i] = Sum{dp[i-nums[j]]}
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

//组合总和(五)
//给定一个由不同整数构成的数组 nums 和一个整数 target ，请你从 nums 找出总和是 target 的组合的个数。
//解集中可以重复使用 nums 中的元素。且解集中数字顺序不同视为不同的组合。
//输入：[1,2,3],4
//返回值：7
//说明：(1,1,1,1) (1,1,2) (1,2,1) (1,3) (2,1,1) (2,2) (3,1)
func CombinationSum5(nums []int, target int) int {
	totalCount := 0
	var bt func(rest int)
	bt = func(rest int) {
		if rest < 0 {
			return
		} else if rest == 0 {
			totalCount++
		}

		for i := 0; i < len(nums); i++ {
			bt(rest - nums[i])
		}
	}
	bt(target)
	return totalCount
}

//电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
//给出数字到字母的映射如下（与电话按键相同）。注意1不对应任何字母。
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
func LetterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	dic := map[int][]byte{
		2: []byte{'a', 'b', 'c'},
		3: []byte{'d', 'e', 'f'},
		4: []byte{'g', 'h', 'i'},
		5: []byte{'j', 'k', 'l'},
		6: []byte{'m', 'n', 'o'},
		7: []byte{'p', 'q', 'r', 's'},
		8: []byte{'t', 'u', 'v'},
		9: []byte{'w', 'x', 'y', 'z'},
	}
	res := make([]string, 0)
	//i表示当前处理的是digits中的第i个数字
	var backtrack func(i int, curr []byte)
	backtrack = func(i int, curr []byte) {
		if i == n {
			res = append(res, string(curr))
			return
		}

		//取数字对应的字母集合
		letter := dic[int(digits[i]-'0')]
		for j := 0; j < len(letter); j++ {
			curr = append(curr, letter[j])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0, []byte{})
	return res
}
