package graph

import "algorithm/container"

// DFS 深度优先遍历，返回遍历Node的Values
func DFS(head *Node) []int {
	if head == nil {
		return nil
	}
	var dfsResult []int
	stack := container.StackBySlice{}
	stack.Push(head)
	dfsResult = append(dfsResult, head.Value)

	visitSet := make(map[*Node]struct{})
	visitSet[head] = struct{}{}

	for !stack.Empty() {
		cur := stack.Pop().(*Node)
		for index := range cur.Nexts {
			if _, ok := visitSet[cur.Nexts[index]]; !ok {
				stack.Push(cur)
				stack.Push(cur.Nexts[index])
				dfsResult = append(dfsResult, cur.Nexts[index].Value)
				visitSet[cur.Nexts[index]] = struct{}{}
				break
			}
		}
	}
	return dfsResult
}
