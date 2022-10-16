package graph

// Edge 图的边
type Edge struct {
	Weight int   // 边的权值
	From   *Node // 边的起点
	To     *Node // 边的终点
}
