package graph

// GenerateGraph 生成图。
// 这里假设图的表达形式如下：
// 	0 1 10
//  1 2 5
//  1 3 8
// 其中，每一行的第一个参数为起始点，第二个参数为终止点，第三个参数为边的权重
func GenerateGraph(input [][]int) *Graph {
	var graph = Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[*Edge]struct{}),
	}
	for _, i := range input {
		fromIndex := i[0]
		toIndex := i[1]
		weight := i[2]
		if _, ok := graph.Nodes[fromIndex]; !ok {
			graph.Nodes[fromIndex] = &Node{
				Value: fromIndex,
			}
		}
		if _, ok := graph.Nodes[toIndex]; !ok {
			graph.Nodes[toIndex] = &Node{
				Value: toIndex,
			}
		}
		edge := Edge{
			Weight: weight,
			From:   graph.Nodes[fromIndex],
			To:     graph.Nodes[toIndex],
		}
		graph.Nodes[fromIndex].Out++
		graph.Nodes[fromIndex].Nexts = append(graph.Nodes[fromIndex].Nexts, graph.Nodes[toIndex])
		graph.Nodes[fromIndex].Edges = append(graph.Nodes[fromIndex].Edges, &edge)

		graph.Nodes[toIndex].In++

		graph.Edges[&edge] = struct{}{}
	}
	return &graph
}
