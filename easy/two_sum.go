package easy

import "fmt"

/*
给出一个整型数组 numbers 和一个目标值 target，请在数组中找出两个加起来等于目标值的数的下标，返回的下标按升序排列。
例如：给出的数组为 [20, 70, 110, 150] , 目标值为90,返回一个数组 [1,2]
要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
*/

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int) //target-numbers[i]为key i为value的哈希表
	res := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		if v, ok := m[numbers[i]]; ok {
			if v < i+1 {
				res[0] = v
				res[1] = i + 1
			} else {
				res[0] = i + 1
				res[1] = v
			}
			break
		} else {
			m[target-numbers[i]] = i + 1
		}
	}
	return res
}

func TestTwoSum() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
