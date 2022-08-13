package sort

func HeapSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	//// 通过heapInsert组成大根堆
	//for i := 0; i < len(array); i++ { // O(N)
	//	array = heapInsert(array, i) // O(logN)
	//}

	// 在给定一数组，通过从下往上heapify，时间复杂度更小：O(N)
	for i := len(array) - 1; i >= 0; i-- { // O(N)
		heapify(array, i, len(array))
	}

	for i := len(array); i > 0; i-- {
		swapByTemp(array, 0, i-1) // 把最大值与当轮堆最后一个值做交换
		heapify(array, 0, i-1)    // 剩下值做heapify
	}
	return array
}

func heapInsert(array []int, index int) []int {
	if index >= len(array) {
		return array
	}
	for {
		fatherNodeIndex := (index - 1) / 2
		if array[index] <= array[fatherNodeIndex] {
			break
		}
		swapByTemp(array, fatherNodeIndex, index)
		index = fatherNodeIndex
	}
	return array
}

func heapify(array []int, index int, heapSize int) []int {
	if heapSize > len(array) || index >= len(array) {
		return array
	}
	for {
		leftChildNodeIndex := index*2 + 1
		//  没左节点，直接退出
		if leftChildNodeIndex >= heapSize {
			break
		}
		largestChildIndex := leftChildNodeIndex
		rightChildNodeIndex := leftChildNodeIndex + 1
		if rightChildNodeIndex < heapSize && array[rightChildNodeIndex] > array[leftChildNodeIndex] {
			largestChildIndex = rightChildNodeIndex
		}
		if array[index] >= array[largestChildIndex] {
			break
		}
		swapByTemp(array, index, largestChildIndex)
		index = largestChildIndex
	}
	return array
}
