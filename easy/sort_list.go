package easy

/*
给定一个节点数为n的无序单链表，对其按升序排序。
要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
*/

func sortInList(head *ListNode) *ListNode {
	//使用一个额外的数组存储链表的所有值，即空间复杂度为O(n)
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	//对这个数组使用归并排序，即时间复杂度为O(nlogn)
	array = MergeSort(array)
	return GenList(array)
}

//对两已排序的数组中的元素进行排序并返回排序后的数组
func mergeArray(left, right []int) []int {
	n1, n2 := len(left), len(right)
	temp := make([]int, n1+n2)
	j, k := 0, 0
	for i := 0; i < n1+n2; i++ {
		for j < n1 && k < n2 {
			if left[j] <= right[k] {
				temp[i] = left[j]
				j++
			} else {
				temp[i] = right[k]
				k++
			}
			i++
		}
		//如果arr1或arr2还没有遍历完，直接添加
		if j < n1 {
			temp[i] = left[j]
			j++
		}
		if k < n2 {
			temp[i] = right[k]
			k++
		}
	}
	return temp
}

func MergeSort(slice []int) []int {
	if slice == nil || len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return mergeArray(MergeSort(slice[0:mid]), MergeSort(slice[mid:]))
}

func TestSortList() {
	head := GenList([]int{-1, 0, -2})
	PrintList(sortInList(head))
}
