package greedyalgorithm

import (
	"container/heap"
	"math"
	"math/rand"
)

// 一块金条切成两半，是需要花费和长度数值一样的铜板的。比如长度为20的金 条，不管切成长度多大的两半，都要花费20个铜板。
// 一群人想整分整块金条，怎么分最省铜板? 例如,给定数组{10,20,30}，代表一共三个人，整块金条长度为10+20+30=60。 金条要分成10,20,30三个部分。 如果先把长度60的金条分成10和50，花费60; 再把长度50的金条分成20和30，花费50;一共花费110铜板。 但是如果先把长度60的金条分成30和30，花费60;再把长度30金条分成10和20， 花费30;一共花费90铜板。
// 输入一个数组，返回分割的最小代价
// 解法：参考哈夫曼算法

type GoldLengthMinHeap []int

func (g GoldLengthMinHeap) Len() int {
	return len(g)
}

func (g GoldLengthMinHeap) Less(i, j int) bool {
	return g[i] < g[j]
}

func (g GoldLengthMinHeap) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g *GoldLengthMinHeap) Push(x interface{}) {
	*g = append(*g, x.(int))
}

func (g *GoldLengthMinHeap) Pop() interface{} {
	res := (*g)[len(*g)-1] // 取最后一个值弹出，由于heap.Pop方法里，会先把头节点跟最后节点做交换，所以这里Pop直接弹出最后一个节点即可
	*g = (*g)[0 : len(*g)-1]
	return res
}

func LessMoneyForSplitGold(input []int) int {
	goldHeap := GoldLengthMinHeap(input)
	heap.Init(&goldHeap)
	var cost int
	for goldHeap.Len() > 1 {
		i := heap.Pop(&goldHeap).(int)
		j := heap.Pop(&goldHeap).(int)
		cost += i + j
		heap.Push(&goldHeap, i+j)
	}
	return cost
}

func LessMoneyForSplitGoldByExhaustive(input []int) int {
	return processForLessMoneyForSplitGold(input, 0)
}

func processForLessMoneyForSplitGold(input []int, pre int) int {
	if len(input) <= 1 {
		return pre
	}
	min := math.MaxInt32
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			min = int(math.Min(float64(min), float64(processForLessMoneyForSplitGold(
				copyGoldLength(input, i, j), pre+input[i]+input[j]))))
		}
	}
	return min
}

func copyGoldLength(input []int, i, j int) []int {
	var res []int
	var newItem int
	for index := range input {
		if index == i || index == j {
			newItem += input[index]
			continue
		}
		res = append(res, input[index])
	}
	res = append(res, newItem)
	return res
}

func generateInput(size int, maxItem int) []int {
	result := make([]int, size)
	for index := range result {
		item := rand.Intn(maxItem + 1)
		result[index] = item
	}
	return result
}
