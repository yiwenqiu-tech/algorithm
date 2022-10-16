package graph

import "algorithm/container"

// TopologySort 拓扑排序。要求有向图，且有入度为0的节点，且没有环
func TopologySort(graph *Graph) []*Node {
	if graph == nil {
		return nil
	}
	var res []*Node
	nodeInMap := make(map[*Node]int)        // 各个node入度Map
	zeroInQueue := container.QueueBySlice{} // 0入度队列
	// 初始化inMap, 并找出0入度节点
	for _, node := range graph.Nodes {
		nodeInMap[node] = node.In
		if node.In == 0 {
			zeroInQueue.Push(node)
		}
	}
	for !zeroInQueue.Empty() {
		n := zeroInQueue.Pop().(*Node)
		res = append(res, n)
		for index := range n.Nexts {
			nodeInMap[n.Nexts[index]]--
			if nodeInMap[n.Nexts[index]] == 0 {
				zeroInQueue.Push(n.Nexts[index])
			}
		}
	}
	return res
}
