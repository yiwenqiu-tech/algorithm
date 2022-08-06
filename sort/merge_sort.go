package sort

// 归并排序--递归方法
func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	process(array, 0, len(array)-1)
	return array
}

func process(array []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	process(array, left, mid)
	process(array, mid+1, right)
	merge(array, left, mid, right)
}

// 归并排序--非递归方法
func mergeSort2(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	mergeSize := 1
	arraySize := len(array)
	for {
		if mergeSize > arraySize {
			break
		}
		left := 0
		for {
			// 若left+mergeSize>=arraySize，即这一轮没有right数，而左侧的数在上一轮已经排序过了，所以可以直接break出来
			if left+mergeSize >= arraySize {
				break
			}
			mid := left + mergeSize - 1
			right := mid + mergeSize
			if right > arraySize-1 {
				right = arraySize - 1
			}
			merge(array, left, mid, right)
			left = right + 1
		}
		// 避免arraySize很大时，导致mergeSize越界，进而影响算法退出
		if mergeSize > arraySize/2 {
			break
		}
		mergeSize = mergeSize << 1
	}
	return array
}

func merge(array []int, left, mid, right int) {
	var temp = make([]int, right-left+1)
	tp := 0
	lp := left
	rp := mid + 1
	for {
		if lp > mid {
			break
		}
		if rp > right {
			break
		}
		if array[lp] <= array[rp] {
			temp[tp] = array[lp]
			lp++
			tp++
		} else {
			temp[tp] = array[rp]
			rp++
			tp++
		}
	}
	if lp <= mid {
		for i := 0; i <= mid-lp; i++ {
			temp[tp+i] = array[lp+i]
		}
	}
	if rp <= right {
		for i := 0; i <= right-rp; i++ {
			temp[tp+i] = array[rp+i]
		}
	}
	for i := 0; i < right-left+1; i++ {
		array[left+i] = temp[i]
	}
}
