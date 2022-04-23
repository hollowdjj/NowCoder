package jz

func FindContinuousSequence(sum int) [][]int {
	if sum == 1 {
		return nil
	}
	limit := 0
	if sum%2 == 1 {
		//例如9可以是4+5
		limit = sum/2 + 1
	} else {
		limit = sum / 2
	}

	left, right := 1, 1
	curr := 0
	res := make([][]int, 0)
	for right <= limit {
		if curr < sum {
			curr += right
			right++
		}
		if curr > sum {
			curr -= left
			left++
		}
		if curr == sum {
			temp := make([]int, 0)
			for i := left; i < right; i++ {
				temp = append(temp, i)
			}
			res = append(res, temp)
			curr -= left
			left++
		}
	}
	return res
}
