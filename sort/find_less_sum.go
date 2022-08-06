package sort

func FindLessSum(array []int) int {
	if len(array) < 2 {
		return 0
	}
	return processForFindLessSum(array, 0, len(array)-1)
}

func processForFindLessSum(array []int, left int, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	return processForFindLessSum(array, left, mid) +
		processForFindLessSum(array, mid+1, right) +
		mergeAndCount(array, left, mid, right)
}

// mergeAndFindReversePair merge的过程中求小和
func mergeAndCount(array []int, left, mid, right int) int {
	count := 0
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
		if array[lp] < array[rp] {
			temp[tp] = array[lp]
			count += array[lp] * (right - rp + 1) // 归并的过程中求小和。由于右侧数据已排好序，所以当左侧比右侧小时，其肯定比右侧往后的数据也小。
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
	return count
}
