package graph

import "container/heap"

type minHeap []*Edge

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight // 最小堆
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*Edge))
}

func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1] // 取最后一个值弹出，由于heap.Pop方法里，会先把头节点跟最后节点做交换，所以这里Pop直接弹出最后一个节点即可
	*h = (*h)[0 : len(*h)-1]
	return res
}

// Prim算法，随机从一个点出发，并把相邻的边放到优先队列里，每次取最邻近的点（注：不能是之前访问过的节点）
func Prim(graph *Graph) []*Edge {
	if graph == nil {
		return nil
	}
	// 小根堆，以边的权重排序
	var edgeHeap = minHeap{}
	var visited = make(map[*Node]struct{})

	var res []*Edge

	for _, value := range graph.Nodes { // for循环主要是兼容上生成森林
		// visited过，直接continue
		if _, ok := visited[value]; ok {
			continue
		}
		visited[value] = struct{}{}
		for _, edge := range value.Edges {
			heap.Push(&edgeHeap, edge)
		}
		for edgeHeap.Len() > 0 {
			minEdge := heap.Pop(&edgeHeap).(*Edge)
			toNode := minEdge.To
			if _, ok := visited[toNode]; !ok {
				res = append(res, minEdge)
				visited[toNode] = struct{}{}
				for _, edge := range toNode.Edges {
					heap.Push(&edgeHeap, edge)
				}
			}
		}
	}
	return res
}
