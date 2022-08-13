package sort

import "math"

// CountSort 计数排序，一种不基于比较的排序，只适用于部分数据类型以及数据状况。
func CountSort(array []int) []int {
	// 获取数组中最大值
	max := math.MinInt32
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	counter := make([]int, max+1)
	for i := 0; i < len(array); i++ {
		counter[array[i]]++
	}
	pointer := 0
	for j := 0; j < len(counter); j++ {
		for {
			if counter[j] <= 0 {
				break
			}
			array[pointer] = j
			pointer++
			counter[j]--
		}
	}
	return array
}
