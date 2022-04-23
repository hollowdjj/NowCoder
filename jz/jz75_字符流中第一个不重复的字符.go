package jz

/*

 */

var dic75 = make(map[byte]int)
var q75 = make([]byte, 0)

func Insert75(ch byte) {
	if _, ok := dic75[ch]; !ok {
		//如果不存在那么加入到队列中
		q75 = append(q75, ch)
	}
	dic75[ch]++
}

func FirstAppearingOnce() byte {
	for len(q75) > 0 && dic75[q75[0]] > 1 {
		//一直pop直到遇到不重复的字符
		q75 = q75[1:]
	}
	if len(q75) == 0 {
		return '#'
	}
	return q75[0]
}
