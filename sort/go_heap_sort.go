package sort

import (
	"container/heap"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	//return h[i] < h[j] // 最小堆
	return h[i] > h[j] // 最大堆
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1] // 取最后一个值弹出，由于heap.Pop方法里，会先把头节点跟最后节点做交换，所以这里Pop直接弹出最后一个节点即可
	*h = (*h)[0 : len(*h)-1]
	return res
}

func GoHeapSort(array []int) []int {
	maxHeap := maxHeap(array)
	heap.Init(&maxHeap)
	for {
		if maxHeap.Len() == 0 {
			return array
		}
		// Pop时，堆顶会换到数组最后一位；所以，若在原数组上调整，最小堆Pop完后按从大到小排序；同理，最大堆Pop后数组是从小到大
		heap.Pop(&maxHeap)
	}
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	//return h[i] < h[j] // 最小堆
	return h[i] < h[j] // 最大堆
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1] // 取最后一个值弹出，由于heap.Pop方法里，会先把头节点跟最后节点做交换，所以这里Pop直接弹出最后一个节点即可
	*h = (*h)[0 : len(*h)-1]
	return res
}

// 已知一个几乎有序的数组，几乎有序是指，如果把数组排好顺序的话，每个元素移动的距离可以不超过k，并且k相对于数组来说比较小，
//  请选择一个合适的排序算法针对这个数据进行排序。
func SortArrayForAlmostSortedArray(array []int, k int) []int {
	minHeap := minHeap{}
	var res []int
	for i := 0; i < len(array); i++ {
		heap.Push(&minHeap, array[i])
		if minHeap.Len() > k {
			res = append(res, heap.Pop(&minHeap).(int))
		}
	}
	for {
		if minHeap.Len() == 0 {
			break
		}
		res = append(res, heap.Pop(&minHeap).(int))
	}
	return res
}
