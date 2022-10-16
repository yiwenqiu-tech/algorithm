package greedyalgorithm

import (
	"container/heap"
	"math"
	"math/rand"
)

//输入:
//正数数组costs
//正数数组profits
//正数k
//正数m
//含义:
//costs[i]表示i号项目的花费 profits[i]表示i号项目在扣除花费之后还能挣到的钱(利润) k表示你只能串行的最多做k个项目
//m表示你初始的资金
//说明: 你每做完一个项目，马上获得的收益，可以支持你去做下一个项目。
//输出: 你最后获得的最大钱数。

type minCostHeap []Progress

func (m minCostHeap) Len() int {
	return len(m)
}

func (m minCostHeap) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func (m minCostHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minCostHeap) Push(x interface{}) {
	*m = append(*m, x.(Progress))
}

func (m *minCostHeap) Pop() interface{} {
	res := (*m)[len(*m)-1] // 取最后一个值弹出，由于heap.Pop方法里，会先把头节点跟最后节点做交换，所以这里Pop直接弹出最后一个节点即可
	*m = (*m)[0 : len(*m)-1]
	return res
}

type maxProfitHeap []Progress

func (m maxProfitHeap) Len() int {
	return len(m)
}

func (m maxProfitHeap) Less(i, j int) bool {
	return m[i].profit > m[j].profit //大顶堆
}

func (m maxProfitHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxProfitHeap) Push(x interface{}) {
	*m = append(*m, x.(Progress))
}

func (m *maxProfitHeap) Pop() interface{} {
	res := (*m)[len(*m)-1] // 取最后一个值弹出，由于heap.Pop方法里，会先把头节点跟最后节点做交换，所以这里Pop直接弹出最后一个节点即可
	*m = (*m)[0 : len(*m)-1]
	return res
}

type Progress struct {
	cost   int
	profit int
}

func initProcess(costs []int, profits []int) []Progress {
	var progress []Progress
	for index := range costs {
		progress = append(progress, Progress{
			cost:   costs[index],
			profit: profits[index],
		})
	}
	return progress
}

func IPO(costs []int, profits []int, k int, m int) int {
	var progress = initProcess(costs, profits)
	// 初始化cost堆
	minHeap := minCostHeap(progress)
	heap.Init(&minHeap)

	cur := m

	maxProfitHeap := maxProfitHeap{}

	for i := 0; i < k; i++ {
		for minHeap.Len() > 0 {
			if cur < minHeap[0].cost { // 第0个节点相当于堆顶
				break
			}
			top := heap.Pop(&minHeap).(Progress)
			heap.Push(&maxProfitHeap, top)
		}
		if maxProfitHeap.Len() == 0 {
			return cur
		}
		maxProfit := heap.Pop(&maxProfitHeap).(Progress)
		cur += maxProfit.profit
	}
	return cur
}

func IPOByExhaustive(costs []int, profits []int, k int, m int) int {
	return processForIPO(costs, profits, k, m)
}

func processForIPO(costs []int, profits []int, k int, m int) int {
	if k < 1 || len(costs) == 0 || len(profits) == 0 || m == 0 {
		return m
	}
	max := m
	for index := range costs {
		if m >= costs[index] {
			max = int(math.Max(float64(max),
				float64(processForIPO(copyExpectIndex(costs, index), copyExpectIndex(profits, index), k-1, m+profits[index]))))
		}
	}
	return max
}

func copyExpectIndex(inputArray []int, inputIndex int) []int {
	var res []int
	for index := range inputArray {
		if index != inputIndex {
			res = append(res, inputArray[index])
		}
	}
	return res
}

func generateInputIPO(size int, maxItem int) ([]int, []int) {
	costs := make([]int, size)
	profits := make([]int, size)
	for index := range costs {
		cost := rand.Intn(maxItem + 1)
		costs[index] = cost
		profit := rand.Intn(maxItem + 1)
		profits[index] = profit
	}
	return costs, profits
}

func generateK(size int) int {
	return rand.Intn(size + 1)
}
