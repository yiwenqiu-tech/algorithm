package xor

// 给定N个数，其中只有一个数出现了奇数次，找出这个奇数
func FindOnlyOneOddNumber(array []int) int {
	oddNum := 0
	for index := range array {
		oddNum ^= array[index]
	}
	return oddNum
}

// 给定N个数，其中只有两个数出现了奇数次，找出这两个奇数
func FindOnlyTwoOddNumber(array []int) (int, int) {
	xorResult := 0
	for index := range array {
		xorResult ^= array[index]
	}
	// 找出异或后结果，从右边往左算的第一个为1的数字。
	resultForFirstNoZeroIndex := xorResult & (^xorResult + 1)
	a := 0
	for index := range array {
		// 根据异或后结果，从右边往左算的第一个为1的数字对原数组进行分组
		if array[index]&resultForFirstNoZeroIndex == 1 {
			a ^= array[index]
		}
	}
	// 找出a之后，a^b^a = (a^a)^b = 0^b = b
	b := xorResult ^ a
	return a, b
}
