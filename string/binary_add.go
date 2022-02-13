package string

/*
二进制求和
输入："101","1"
输出："110"
*/

func BinaryAdd(A string, B string) string {
	i, j := len(A)-1, len(B)-1
	carry := 0
	res := make([]byte, max(len(A), len(B))+1)
	k := len(res) - 1
	for i >= 0 && j >= 0 {
		val := int(A[i]-'0') + int(B[j]-'0') + carry
		carry = val / 2
		val %= 2
		if val == 0 {
			res[k] = '0'
		} else {
			res[k] = '1'
		}
		i--
		j--
		k--
	}
	for i >= 0 {
		val := int(A[i]-'0') + carry
		carry = val / 2
		val %= 2
		if val == 0 {
			res[k] = '0'
		} else {
			res[k] = '1'
		}
		i--
		k--
	}
	for j >= 0 {
		val := int(B[j]-'0') + carry
		carry = val / 2
		val %= 2
		if val == 0 {
			res[k] = '0'
		} else {
			res[k] = '1'
		}
		j--
		k--
	}
	if carry == 0 {
		return string(res[1:])
	}
	res[0] = '1'
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
