package codetop

/*
森林中有未知数量的兔子。提问其中若干只兔子 "还有多少只兔子与你（指被提问的兔子）颜色相同?"
将答案收集到一个整数数组 answers 中，其中 answers[i] 是第 i 只兔子的回答。
给你数组 answers ，返回森林中兔子的最少数量。

输入：answers = [1,1,2]
输出：5
解释：
两只回答了 "1" 的兔子可能有相同的颜色，设为红色。
之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
设回答了 "2" 的兔子为蓝色。
此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。
因此森林中兔子的最少数量是 5 只：3 只回答的和 2 只没有回答的。
*/

func numRabbits(answers []int) int {
	//1. 报不同数字的兔子一定不是同一个颜色的
	//2. 有兔子报了n，那么兔子的最少数量要加n+1
	//3. 有兔子报了n，要想使得兔子数量最少，需忽略后续n个报n的兔子。
	//例如，[2,2,2]，兔子数量最少的情况三个兔子颜色一样。第一个报2的加了3，后面2个2就要忽略
	res := 0
	dic := make(map[int]int) //数量n以及当前已经被忽略的次数
	for _, count := range answers {
		if count == 0 {
			res++
			continue
		}
		if v, ok := dic[count]; !ok || v == count {
			res += count + 1
			dic[count] = 0
		} else {
			dic[count]++
		}
	}
	return res
}
