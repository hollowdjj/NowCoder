package array

/*
数组排序

数组排序算法中有一个重要的概念叫做“算法的稳定性”。排序算法的稳定性指的是，假定在待排序的记录序列中，存在多个
具有相同的关键字的记录，若经过排序，这些记录的相对次序保持不变，即在原序列中，a[i]=a[j]，且a[i]在a[j]之前，而
在排序后的序列中，a[i]仍在a[j]之前，则称这种排序算法是稳定的；否则称为不稳定的。
*/

//插入排序，时间复杂度为O(n^2)，空间间复杂度为O(1)，是一种稳定排序算法
func InsertSort(arr []int) []int {
	//即每次固定一个数(arr[i])，然后一直往前交换，直到找到正确的位置。例如[3,2,1]
	//i = 0; [3,2,1]
	//i = 1; [2,3,1]
	//i = 2; [1,2,3]
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j >= 1; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

//冒泡排序，最优O(n)，最坏O(n^2)，是一种稳定排序算法
func BubbleSort(arr []int) []int {
	//从左到右依次比较arr[j]和arr[j-1]的大小，并将大的那个交换到右边。一趟交换完成后，数组中的最后一个
	//元素即为数组的最大值。第二趟交换完成后，倒数第二个元素为数组倒数第二大的值。因此，n个元素需要进行
	//n-1趟交换。而第i趟交换又要进行n-i次左右交换。因此，冒泡排序的最优时间复杂度为O(n)，即只需要一趟
	//交换即可。最坏情况下的时间复杂度为O(n^2)，即数组本身是逆序的(4,3,2,1,0)。
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

//选择排序，时间复杂度O(n^2)，空间复杂度O(1)，是一种不稳定排序算法([5,8,5,2,9])
func SelectSort(arr []int) []int {
	//从arr[i+1....n]中选取最小的那个值，与arr[i]交换。
	n := len(arr)
	for i := 0; i < n; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}

//快排，填坑式做法。最优O(nlogn)，最坏O(n^2)
func QuickSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	//需要注意的是，我们必须选择数组的第一个数作为pivot，而不能选择其他的数。一些改进方法，例如随机选取
	//数组中的一个数做为pivot时，也只是将那个随机选取的数和第一个数交换，pivot始终是指数组的头元素。
	pivot := arr[0]
	left, right := 0, len(arr)-1
	for left < right {
		//首先从右往左找到第一个小于pivot的数。然后将其赋值给arr[left](之前挖好的坑)
		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			left++
		}
		//然后从左往右找到第一个大于pivot的数。然后将其赋值给arr[right](之前挖好的坑)
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right--
		}
	}
	arr[left] = pivot
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}

//归并排序，时间复杂度O(nlogn)，空间复杂度O(n)。归并排序是一种稳定的排序算法，但缺点在于无法原地排序
func MergeSort(arr []int) []int {
	var m func(arr []int, left, right int)
	m = func(arr []int, left, right int) {
		if left < right {
			mid := (right + left) / 2
			m(arr, left, mid)
			m(arr, mid+1, right)
			doMerge(left, mid, right, arr)
		}
	}
	m(arr, 0, len(arr)-1)
	return arr
}

func doMerge(left, mid, right int, arr []int) {
	//合并两个相邻且已排序的数组，arr[left,mid]和arr[mid+1,right]
	temp := make([]int, right-left+1)
	i, j := left, mid+1
	k := 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	//赋值
	k-- //这里必须减1，因为此时k = len(temp)
	for m := right; m >= left; m-- {
		arr[m] = temp[k]
		k--
	}
}

//堆排序，时间复杂度O(nlogn)，空间复杂度O(1)，是不稳定排序算法。
func HeapSort(arr []int) []int {
	//堆是具有以下性质的完全二叉树：每个结点的值都大于或等于其左右孩子结点的
	//值，称为大顶堆；或者每个结点的值都小于或等于其左右孩子结点的值，称为小顶堆。
	//如下图:
	//          大顶堆						小顶堆
	//			  50						  10
	//       45        40				20			15
	//   20      25 35     30       25       50 30       40
	//10     15				    35	    45
	//我们将大顶堆按照层序遍历映射到数组中即为[50,45,40,20,25,35,30,10,15]。因此,
	//arr[i]的左孩子为arr[2i+1]右孩子为arr[2i+2]。
	//对于一个大顶堆数组，其排序过程如下：
	//1.将数组的头元素与尾元素交换。然后对前面的元素重新构建大顶堆
	//2.将数字头元素与倒数第二个元素交换，然后对前面的元素重新构建大顶堆
	//......

	//首先将无序数组构建为大顶堆。从最后一个非叶子节点开始向前遍历。需要注意的是，完全二叉树
	//的最后一个非叶子节点的索引为len(arr)/2 - 1。其推导过程如下：
	//1. 假设最后一个非叶子节点只有一个左孩子，那么该左孩子的索引为n-1。n-1 = 2*i+1 -->i = n/2 - 1
	//2. 假设最后一个非叶子节点有两个孩子。那么n-1 = 2*i + 2; n-2 = 2*i + 1 ---> i = (n-1)/2 -1
	//最后一个非叶子节点有两个孩子时，n一定是奇数，(n-1)/2 == n/2
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		AdjustMaxHeap(arr, i, n)
	}
	//排序
	for i := n - 1; i > 0; i-- {
		//交换arr[0]和arr[i]，然后将arr[0...i]调整为大顶堆。这里导致了堆排序不是一个稳定算法。例如[5,4,4,3,4]
		arr[i], arr[0] = arr[0], arr[i]
		AdjustMaxHeap(arr, 0, i)
	}

	return arr
}

//将以parentIndex为根节点的二叉树调整为大顶堆
func AdjustMaxHeap(arr []int, parentIndex int, length int) {
	childIndex := parentIndex*2 + 1 //左孩子
	for childIndex < length {
		//选取左右孩子中的较大值
		if childIndex+1 < length && arr[childIndex+1] > arr[childIndex] {
			childIndex++
		}
		//判断根节点是否大于等于其左右孩子
		if arr[parentIndex] >= arr[childIndex] {
			break
		}
		//将较大值交换到根节点，然后继续从发生了交换的子节点往下判断
		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
}
