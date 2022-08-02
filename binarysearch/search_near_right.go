package binarysearch

func SearchNearRight(array []int, num int) int {
	if len(array) == 0 {
		return -1
	}
	left := 0
	right := len(array) - 1
	index := -1
	for {
		if left > right {
			break
		}
		mid := left + (right-left)/2
		if array[mid] <= num {
			index = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}
