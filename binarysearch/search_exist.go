package binarysearch

// SearchExist 利用二分查找，寻找数值是否存在。注意：array必须有序
func SearchExist(array []int, num int) bool {
	// 避免越界
	if len(array) == 0 {
		return false
	}
	left := 0
	right := len(array) - 1
	for {
		if left == right {
			break
		}
		mid := left + (right-left)/2
		if array[mid] == num {
			return true
		}
		if array[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return array[left] == num
}
