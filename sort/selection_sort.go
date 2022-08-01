package sort

// SelectionSort 选择排序：每轮把最小的找出来放在前面
// 第一轮：遍历0-N-1，把最小的放在第0位
// 第二轮：遍历1-N-1，把1~N-1位最小值放在第1位
// 以此类推
func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		if i != minIndex {
			swapByXOR(array, i, minIndex)
		}
	}
	return array
}
