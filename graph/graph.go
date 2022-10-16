package graph

// Graph 图的表达
type Graph struct {
	Nodes map[int]*Node      // 点集
	Edges map[*Edge]struct{} // 边集
}
