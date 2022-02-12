package string

/*
大数乘法
*/

func BigNumMultiply(s string, t string) string {
	/*
			        9  8		  s
			×       2  1	      t
			-------------
			       (9)(8)  <---- 第1趟: 98×1的每一位结果
			  (18)(16)     <---- 第2趟: 98×2的每一位结果
			-------------
			  (18)(25)(8)  <---- 这里就是相对位的和，还没有累加进位
		s[i]*s[j]的结果是放在res[i+j+1]的(+1是因为res[0]是保留的进位)。然后将res从后往前
		相加
	*/
	numS, numT := []byte(s), []byte(t)
	tempRes := make([]int, len(s)+len(t))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			tempRes[i+j+1] += int(numS[i]-'0') * int(numT[j]-'0')
		}
	}

	//处理进位
	for i := len(tempRes) - 1; i > 0; i-- {
		if tempRes[i] >= 10 {
			tempRes[i-1] += tempRes[i] / 10
			tempRes[i] %= 10
		}
	}
	//去除前导0
	k := 0
	for ; k < len(tempRes); k++ {
		if tempRes[k] != 0 {
			break
		}
	}
	if k == len(tempRes) {
		return "0"
	}
	tempRes = tempRes[k:]
	res := make([]rune, len(tempRes))
	for i, n := range tempRes {
		res[i] = rune(n + '0')
	}

	return string(res)
}
