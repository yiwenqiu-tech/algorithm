package sort

// 荷兰国旗问题简化版
// 把小于等于某值的放在左侧，其余放在右侧
func DutchNationalFlagSimplifiedVersion(array []int, num int) []int {
	less := -1
	for i := 0; i < len(array); i++ {
		if array[i] <= num {
			less++
			swapByTemp(array, i, less)
		}
	}
	return array
}

// 荷兰🇳🇱国旗问题
//	从国旗标识可以看到，荷兰国旗问题其实就是分类问题，例如：
//	给定一个数组以及一个数字，要求把小于给定数字的放在数组左侧，等于的放于中间，大于的放在右侧
func DutchNationalFlag(array []int, num int) []int {
	less := -1
	more := len(array)
	point := 0
	for {
		if point == more {
			break
		}
		if array[point] < num {
			less++
			swapByTemp(array, point, less)
			point++
		} else if array[point] == num {
			point++
		} else {
			more--
			swapByTemp(array, point, more)
		}
	}
	return array
}
