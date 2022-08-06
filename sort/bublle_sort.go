package sort

import "sort"

// BubbleSort 冒泡排序：每一轮把最大的数值往最后挪
// 第一轮：遍历0-N-1，把最大的放在第N-1位
// 第二轮：遍历0-N-2，把1～N-2位最大值放在第N-2位
// 以此类推
func BubbleSort(array []int) []int {
	for i := len(array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				swapByXOR(array, j, j+1)
			}
		}
	}
	return array
}

func golangSort(array []int) []int {
	sort.Sort(sort.IntSlice(array))
	return array
}
