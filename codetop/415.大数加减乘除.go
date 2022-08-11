package codetop

import (
	"fmt"
	"strings"
)

/*
使用字符串进行大数的加减乘除
*/

//大数加法，逆向遍历注意进位即可
func addString(num1, num2 string) string {
	s1, s2 := []byte(num1), []byte(num2)
	res := make([]string, max(len(s1), len(s2))+1)
	i, i1, i2 := len(res)-1, len(num1)-1, len(num2)-1
	carry := 0
	for i1 >= 0 || i2 >= 0 {
		sum := carry
		if i1 >= 0 {
			sum += int(num1[i1] - '0')
		}
		if i2 >= 0 {
			sum += int(num2[i2] - '0')
		}
		res[i] = fmt.Sprintf("%d", sum%10)
		carry = sum / 10
		i--
		i1--
		i2--
	}
	if carry != 0 {
		res[i] = fmt.Sprintf("%d", carry)
	}
	return strings.Join(res, "")
}

//大数减法，减法需要借位，并且当是小减大时，会出现负数。因此，为了编程方便
//使用大减小。
func SubString(num1, num2 string) string {
	sign := 1
	bigger, smaller := "", ""
	if len(num1) == len(num2) {
		if num1 > num2 {
			bigger, smaller = num1, num2
		} else {
			sign = -1
			bigger, smaller = num2, num1
		}
	} else {
		if len(num1) > len(num2) {
			bigger, smaller = num1, num2
		} else {
			sign = -1
			bigger, smaller = num2, num1
		}
	}

	borrow := 0
	res := make([]string, len(bigger))
	i, ib, is := len(bigger)-1, len(bigger)-1, len(smaller)-1
	for is >= 0 {
		sub := int(bigger[ib]-'0') - int(smaller[is]-'0') - borrow
		if sub < 0 {
			borrow = 1
			sub += 10
		} else {
			borrow = 0
		}
		res[i] = fmt.Sprintf("%d", sub%10)
		i--
		ib--
		is--
	}
	for ib >= 0 {
		sub := int(bigger[ib]-'0') - borrow
		if sub < 0 {
			borrow = 1
			sub += 10
		} else {
			borrow = 0
		}
		res[i] = fmt.Sprintf("%d", sub%10)
		i--
		ib--
	}

	ret := ""
	if sign == -1 {
		ret = "-"
	}
	//减法会出现前导0，所以要去除前导0
	ret += strings.TrimLeft(strings.Join(res, ""), "0")

	return ret
}

//大数乘法。让我们首先考虑多位数乘一位数的情况，我们直接将和每一位数相乘的结果保存到数组中。
//  9   9   9
//          9
//  81  81  81
//然后，在遍历数组进行进位，最终即可得到8991。对于多位数乘多位数的情况：
//  	9     9
//  	9     9
// 		81    81
//	81  81
//       =
// 81   162   81，遍历数组，进位即可得到9801。
//两数字相乘后的结果的长度最长为len(num1)+len(num2)。需要注意的是，我们必须逆序遍历num1和num2，否则
//会出现2*3=60的情况。逆序遍历时，num1的第i位，乘num2的第j位的结果应该加到nums数组中的i+j位。
func MulString(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	nums := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			index := i + j + 1
			nums[index] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}

	n := len(nums)
	carry := 0
	for i := n - 1; i >= 0; i-- {
		sum := nums[i] + carry
		nums[i] = sum % 10
		carry = sum / 10
	}
	if carry != 0 {
		nums[0] = carry
	}
	res := make([]string, n)
	for i := range nums {
		res[i] = fmt.Sprintf("%d", nums[i])
	}
	return strings.TrimLeft(strings.Join(res, ""), "0")
}

//大数除法。对于除法而言，可以使用加法求解，但当数很大时，会较为耗时。
//求a/b，其实就是求a中有多少个b。使用二分法求解，即a > b * 2^k时，res += 2^k
func DivStrign(num1, num2 string) string {
	return ""
}
