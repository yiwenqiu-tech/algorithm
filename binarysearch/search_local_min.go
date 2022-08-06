package binarysearch

// SearchLocalMin 寻找局部最小值；给定一个无序数组，且无序数组中相邻的数一定不相等，寻找一个局部最小值。
// 局部最小值定义：
// 1.在位置0上，若1位置比他大，则他是局部最小值；
// 2.在位置N上，若N-1位置比他大，则他是局部最小值；
// 3.若i位置中，i-1、i+1均比他大，则他是局部最小值
func SearchLocalMin(array []int) int {
	if len(array) <= 1 {
		return -1
	}
	if array[0] < array[1] {
		return 0
	}
	if array[len(array)-1] < array[len(array)-2] {
		return len(array) - 1
	}
	// 注意初始值，避免越界
	left := 1
	right := len(array) - 2
	for {
		if left > right {
			break
		}
		mid := left + (right-left)/2
		if array[mid] > array[mid-1] {
			right = mid - 1
		} else if array[mid] > array[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
