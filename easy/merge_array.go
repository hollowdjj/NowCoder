package easy

import "fmt"

/*
输入：	[4,5,6],[1,2,3]
返回值： [1,2,3,4,5,6]
说明：   A数组为[4,5,6]，B数组为[1,2,3]，后台程序会预先将A扩容为[4,5,6,0,0,0]，B还是为[1,2,3]，
        m=3，n=3，传入到函数merge里面，然后请同学完成merge函数，将B的数据合并A里面，最后后台程序输出A数组
*/

func Merge(A []int, m int, B []int, n int) {
	a, b := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		//当数组A还没有走完且A[a] > B[b]或者数组B以及走完了时
		if b < 0 || (a >= 0 && A[a] > B[b]) {
			A[i] = A[a]
			a--
		} else {
			A[i] = B[b]
			b--
		}
	}

	fmt.Println(A)
}

func TestMerge() {
	A := []int{4, 7, 9, 0, 0, 0}
	B := []int{3, 5, 8}
	Merge(A, 3, B, 3)

}
