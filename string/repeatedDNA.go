package string

/*
重复的DNA序列
所有的 DNA 序列都是由 'A' , ‘C’ , 'G' , 'T' 字符串组成的，例如 'ACTGGGC' 。
请你实现一个函数找出所有的目标子串，目标连续子串的定义是，长度等于 10 ，且在 DNA 序列
中出现次数超过 1 次的连续子串（允许两个连续子串有重合的部分，如下面的示例2所示）。
（注：返回的所有目标子串的顺序必须与原DNA序列的顺序一致，如下面的示例1所示）

输入："AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出： ["AAAAACCCCC","CCCCCAAAAA"]
说明： "AAAAACCCCC"和"CCCCCAAAAA"长度等于 10 且在DNA序列中分别出现了 2 次。
	   不能返回["CCCCCAAAAA","AAAAACCCCC"]，因为在原DNA序列中，"AAAAACCCCC"要比"CCCCCAAAAA"先出现。
*/

func RepeatedDNA(DNA string) []string {
	n := len(DNA)
	if n < 11 {
		return nil
	}

	dic := make(map[string]int) //记录子串出现的次数
	//滑动窗口
	for i := 9; i < n; i++ {
		substr := DNA[i-9 : i+1]
		dic[substr]++
	}
	//由于要按照重复子串第一次出现的顺序输出，这里再遍历一次
	res := make([]string, 0)
	for i := 9; i < n; i++ {
		substr := DNA[i-9 : i+1]
		if v, ok := dic[substr]; ok && v > 1 {
			res = append(res, substr)
		}
		delete(dic, substr) //因为有重复子串，所以要删除
	}
	return res
}
