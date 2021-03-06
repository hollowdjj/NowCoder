package jz

/*
牛客最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。同事Cat对Fish写的内容颇感兴趣，有一天他向
Fish借来翻看，但却读不懂它的意思。例如，“nowcoder. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，正确的句子
应该是“I am a nowcoder.”。Cat对一一的翻转这些单词顺序可不在行，你能帮助他么？

输入："nowcoder. a am I"
输出："I am a nowcoder."
*/

func ReverseSentence(str string) string {
	res := ""
	n := len(str)
	left, right := n, n-1
	for right >= 0 {
		if str[right] == ' ' {
			res += str[right+1 : left]
			res += " "
			left = right
		}
		right--
	}

	res += str[right+1 : left]
	return res
}
