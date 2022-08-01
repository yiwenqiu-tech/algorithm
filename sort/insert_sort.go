package sort

// InsertSort 插入排序（类似于玩扑克牌，每一次拿一张然后跟之前已经拿的比较，放到合适的位置）
func InsertSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				swapByXOR(array, j, j-1)
			} else {
				break
			}
		}
	}
	return array
}
