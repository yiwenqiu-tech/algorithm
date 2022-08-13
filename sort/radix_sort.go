package sort

import "math"

// RadixSort 基数排序
func RadixSort(array []int, radix int) []int {
	if len(array) < 2 {
		return array
	}
	return radixSort(array, 0, len(array)-1, maxBits(array, radix), radix)
}

func maxBits(array []int, radix int) int {
	max := math.MinInt32
	for _, a := range array {
		if a > max {
			max = a
		}
	}
	res := 0
	for {
		if max == 0 {
			break
		}
		max = max / radix
		res++
	}
	return res
}

func radixSort(array []int, L int, R int, maxBit int, radix int) []int {
	temp := make([]int, R-L+1)
	for i := 1; i <= maxBit; i++ {
		count := make([]int, radix)
		for j := L; j <= R; j++ {
			digit := getDigit(array[j], i, radix)
			count[digit]++
		}
		for j := 1; j < len(count); j++ {
			count[j] = count[j] + count[j-1]
		}
		for j := R; j >= L; j-- {
			digit := getDigit(array[j], i, radix)
			temp[count[digit]-1] = array[j]
			count[digit]--
		}
		for j := L; j <= R; j++ {
			array[j] = temp[j-L]
		}
	}
	return array
}

func getDigit(num int, d int, radix int) int {
	return (num / int(math.Pow(float64(radix), float64(d-1)))) % 10
}
