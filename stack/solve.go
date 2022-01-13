package stack

import (
	"nowcoder/utility"
	"unicode"
)

/*
表达式求值：
请写一个整数计算器，支持加减乘三种运算和括号。

数据范围：0≤∣s∣≤100，保证计算结果始终在整型范围内
要求：空间复杂度：O(n)，时间复杂度 O(n)
*/

func getInt(i *interface{}) int {
	return (*i).(int)
}

func Solve(s string) int {
	var nums utility.Stack //存放表达式中的数字、括号中子表达式的结果、乘法以及除法的结果。
	number := 0            //当前数字
	sign := '+'            //上一个符号
	n := len(s)
	for i := 0; i < n; i++ {
		v := rune(s[i])
		if v == ' ' {
			continue
		}
		if unicode.IsDigit(v) {
			//如果v是数字，那么依靠进位递归得到完整的数字，例如“324”
			number = number*10 + int(v-'0')
		}
		if v == '(' {
			//如果v是左括号，那么一直往前找到匹配的右括号。然后递归计算括号里面的子表达式。
			//因为，一对括号中还可能存在其他括号对，因此需要在遇到'('时加1，遇到')'时减一
			//当减到0时，对于的右括号就是与之匹配的右括号。
			count := 1
			j := i + 1
			for count > 0 {
				if s[j] == '(' {
					count++
				} else if s[j] == ')' {
					count--
				}
				j++
			}
			number = Solve(s[i+1 : j-1]) //递归计算括号内表达式的值，赋值给number。
			i = j - 1                    //匹配的右括号的索引为j-1
		}
		if !unicode.IsDigit(v) || i == n-1 {
			//如果v是运算符的话，因为优先级的关系需要判断一下上一个运算符是什么。如果上一个运算符是
			//'+'，那么直接push当前数字；如果是'-'，那么push当前数字的相反数;如果是'*'，那么将当前数
			//和前一个数相乘，并push相乘的结果；除法和乘法同理。 (2*3 + 1)。
			//这里或条件中的i==n-1是为了表达式的最后一个元素为数字的情况
			switch sign {
			case '+':
				nums.Push(number)
			case '-':
				nums.Push(-1 * number)
			case '*':
				nums.Push(getInt(nums.Pop()) * number)
			case '/':
				nums.Push(getInt(nums.Pop()) / number)
			}
			sign = v
			number = 0
		}
	}
	//现在只剩下加法，把nums中的所有数字都加起来即可
	sum := 0
	for !nums.Empty() {
		sum += getInt(nums.Pop())
	}
	return sum
}
