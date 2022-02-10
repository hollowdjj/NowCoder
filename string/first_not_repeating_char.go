package string

/*
找到字符串中第一个只出现一次的元素
*/

func FirstNotRepeatingChar(str string) int {
	dic := make(map[rune]int)
	for i, b := range str {
		if _, ok := dic[b]; ok {
			delete(dic, b)
		} else {
			dic[b] = i
		}
	}

	res := 10001
	for _, v := range dic {
		if v < res {
			res = v
		}
	}
	if res == 10001 {
		return -1
	}
	return res
}
