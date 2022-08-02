package binarysearch

// SearchNearLeft 寻找大于等于某个数最左侧的数
func SearchNearLeft(array []int, num int) int {
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
		if array[mid] >= num { // 因为找的是>=某个数最左侧的数
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}
