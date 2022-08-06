package sort

// FindReversePair2 寻找逆序对，暴力解法。这里应用于作对数器
func FindReversePair2(array []int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				count++
			}
		}
	}
	return count
}

// FindReversePair 寻找逆序对，递归法
func FindReversePair(array []int) int {
	if len(array) < 2 {
		return 0
	}
	return processForFindReversePair(array, 0, len(array)-1)
}

func processForFindReversePair(array []int, left int, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	return processForFindReversePair(array, left, mid) +
		processForFindReversePair(array, mid+1, right) +
		mergeAndFindReversePair(array, left, mid, right)
}

// mergeAndFindReversePair merge的过程中找出逆序对
func mergeAndFindReversePair(array []int, left, mid, right int) int {
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
		if array[lp] <= array[rp] {
			temp[tp] = array[lp]
			lp++
			tp++
		} else {
			temp[tp] = array[rp]
			rp++
			tp++
			count += mid - lp + 1
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
