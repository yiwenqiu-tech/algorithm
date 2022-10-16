package graph

// Node 图的顶点
type Node struct {
	Value int     // 顶点的值
	In    int     // 顶点的入度
	Out   int     // 顶点的出度
	Nexts []*Node // 顶点的Next点集
	Edges []*Edge // 顶点的相关边集，从当前点出发的
}
