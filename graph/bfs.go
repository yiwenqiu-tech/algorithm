package graph

import "algorithm/container"

// BFS 宽度优先遍历, 返回宽度优先遍历的Node的Value
func BFS(head *Node) []int {
	if head == nil {
		return nil
	}
	var bfsResult []int
	queue := container.QueueBySlice{}
	queue.Push(head)
	// 记录遍历过的节点
	nodeSet := make(map[*Node]struct{})
	nodeSet[head] = struct{}{}
	for !queue.Empty() {
		node := queue.Pop().(*Node)
		bfsResult = append(bfsResult, node.Value)
		for index := range node.Nexts {
			if _, ok := nodeSet[node.Nexts[index]]; !ok {
				queue.Push(node.Nexts[index])
				nodeSet[node.Nexts[index]] = struct{}{}
			}
		}
	}
	return bfsResult
}
