package jz_II

func AddBinary(a string, b string) string {
	res := make([]rune, max(len(a), len(b))+1)
	res[0] = '0'
	i, j, k, carry := len(a)-1, len(b)-1, len(res)-1, 0
	dic := map[int]rune{1: '1', 0: '0'}
	for i >= 0 || j >= 0 || carry > 0 {
		ni, nj := 0, 0
		if i >= 0 {
			ni = int(a[i] - '0')
		}
		if j >= 0 {
			nj = int(b[j] - '0')
		}
		sum := ni + nj + carry
		carry = sum / 2
		res[k] = dic[sum%2]
		i--
		j--
		k--
	}

	if res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
