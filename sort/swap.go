package sort

func swapByXOR(array []int, i, j int) {
	array[i] = array[i] ^ array[j]
	array[j] = array[i] ^ array[j]
	array[i] = array[i] ^ array[j]
}

func swapByTemp(array []int, i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
