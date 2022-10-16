package graph

import "math"

// Dijkstra(狄杰斯特拉)算法，主要解决从一个顶点出发到其余各顶点的最短路径算法，解决的是有权图中的最短路径问题
// 	算法要求：图中不存在有负权的边（或者严谨的说不存在负权的环，因为这样计算不出来最短路径）
//	算法思想：
//		1.通过Dijkstra计算图G中的最短路径时，需要指定起点s(即从顶点s开始计算)。
//		2.此外，引进两个集合S和U。S的作用是记录已求出最短路径的顶点(以及相应的最短路径长度)，而U则是记录还未求出最短路径的顶点(以及该顶点到起点s的距离)。
//		3.初始时，S中只有起点s；U中是除s之外的顶点，并且U中顶点的路径是”起点s到该顶点的路径”。然后，从U中找出路径最短的顶点，并将其加入到S中；接着，更新U中的顶点和顶点对应的路径。 然后，再从U中找出路径最短的顶点，并将其加入到S中；接着，更新U中的顶点和顶点对应的路径。 … 重复该操作，直到遍历完所有顶点。
func Dijkstra(from *Node) map[*Node]int {
	distinctMap := make(map[*Node]int)
	distinctMap[from] = 0
	selectedNode := make(map[*Node]struct{}) // 记录已求出最短路径的顶点
	node := findMinLengthNodeAndNotInSelectedNode(distinctMap, selectedNode)
	for node != nil {
		// 当前node距离from的距离
		var fromToNodeDistinct = distinctMap[node]
		for index := range node.Edges {
			edge := node.Edges[index]
			toNode := edge.To
			curLength, exist := distinctMap[toNode]
			if !exist {
				// distinctMap不存在，证明之前没有路径可达，所以直接添加
				distinctMap[toNode] = fromToNodeDistinct + edge.Weight
			} else {
				// 原来存在路径的，需要比较两条路径的大小，若新路径更小则直接更新
				if fromToNodeDistinct+edge.Weight < curLength {
					distinctMap[toNode] = fromToNodeDistinct + edge.Weight
				}
			}
		}
		selectedNode[node] = struct{}{}
		node = findMinLengthNodeAndNotInSelectedNode(distinctMap, selectedNode)
	}
	return distinctMap
}

// findMinLengthNodeAndNotInSelectedNode 找出当前有最小长度且不在已明确最短距离范围内的顶点
func findMinLengthNodeAndNotInSelectedNode(distinctMap map[*Node]int, selectedNode map[*Node]struct{}) *Node {
	var res *Node
	min := math.MaxInt32
	for node, length := range distinctMap {
		if _, ok := selectedNode[node]; !ok && length < min {
			res = node
			min = length
		}
	}
	return res
}
