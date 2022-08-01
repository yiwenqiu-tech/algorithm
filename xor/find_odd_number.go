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
func FindOnlyTwoOddNumber(array []int) int {
	//xorResult := 0
	//for index := range array {
	//	xorResult ^= array[index]
	//}
	//firstNoZeroIndex := 0
	//for {
	//	if (xorResult >> firstNoZeroIndex) & 1 == 0 {
	//		firstNoZeroIndex++
	//	}else {
	//		break
	//	}
	//}
	//for index := range array {
	//	if array[index] >> firstNoZeroIndex {
	//
	//	}
	//}
	//return oddNum
}
