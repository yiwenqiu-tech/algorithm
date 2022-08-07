package sort

func QuickSortV1(array []int) []int {
	if len(array) < 2 {
		return array
	}
	processV1(array, 0, len(array)-1)
	return array
}

func processV1(array []int, left int, right int) {
	if right-left < 2 { // 剩下一个数时，则不需要排序
		return
	}
	mid := partition1(array, left, right)
	processV1(array, left, mid-1) // 由于num必须存在于数组内，所以其实排完后mid的位置是排好的
	processV1(array, mid+1, right)
}

// 根据DutchNationalFlagSimplifiedVersion题解的思路解，即小于等于放一边，大于的放一边
func partition1(array []int, left int, right int) int {
	less := left - 1
	for i := left; i <= right; i++ {
		if array[i] <= array[right] {
			less++
			swapByTemp(array, i, less)
		}
	}
	return less
}

func QuickSortV2(array []int) []int {
	if len(array) < 2 {
		return array
	}
	processV2(array, 0, len(array)-1)
	return array
}

func processV2(array []int, left int, right int) {
	if right-left < 2 { // 剩下一个数时，则不需要排序
		return
	}
	pLeft, pRight := partition2(array, left, right)
	processV2(array, left, pLeft)
	processV2(array, pRight, right)
}

// 根据DutchNationalFlag题解的思路解，即小于放一边，等于放中间，大于的放一边
func partition2(array []int, left int, right int) (int, int) {
	num := array[right]
	less := left - 1
	more := right
	point := left
	for {
		if point >= more {
			break
		}
		if array[point] < num {
			less++
			swapByTemp(array, point, less)
			point++
		} else if array[point] == num {
			point++
		} else {
			more--
			swapByTemp(array, point, more)
		}
	}
	swapByTemp(array, more, right)
	more++            // 交换后，大于数字里往右挪一格
	return less, more //小于数字里的最右侧，大于数字里的最左侧
}

func QuickSortV3(array []int) []int {
	if len(array) < 2 {
		return array
	}
	processV3(array, 0, len(array)-1)
	return array
}

func processV3(array []int, left int, right int) {
	if right-left < 2 { // 剩下一个数时，则不需要排序
		return
	}
	pLeft, pRight := partition2(array, left, right)
	processV3(array, left, pLeft)
	processV3(array, pRight, right)
}
