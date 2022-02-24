package array

/*
找到第K大的数
*/
func FindKth(a []int, n int, K int) int {
	//使用快排加二分法找到第K大的数
	return quickSort(a, K)
}

func quickSort(arr []int, k int) int {
	low, high := 0, len(arr)-1
	pivot := arr[0]
	for low < high {
		//从右往左找到第一个小于pivot的数，赋值给low
		for ; low < high; high-- {
			if arr[high] < pivot {
				arr[low] = arr[high]
				low++
				break
			}
		}
		//从左往右找到第一个大于pivot的数，赋值给high
		for ; low < high; low++ {
			if arr[low] > pivot {
				arr[high] = arr[low]
				high--
				break
			}
		}
	}
	arr[low] = pivot
	if len(arr)-low == k {
		//找到了第k大的元素
		return pivot
	}
	if len(arr)-low > k {
		return quickSort(arr[low+1:], k)
	}

	//此时右数组中有len(arr)-low个大数，还需要k-len(arr)+low个数
	return quickSort(arr[:low], k+low-len(arr))
}
