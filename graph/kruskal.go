package graph

import "sort"

// 克鲁斯卡尔(Kruskal)算法
//  以边为目标去构建，因为权值是在边上，直接去找最小权值的边来构建生成，在构建时要考虑是否会形成环路而已

type Edges []*Edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Less(i, j int) bool {
	return e[i].Weight < e[j].Weight
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func Kruskal(graph *Graph) []*Edge {
	var e Edges
	for k := range graph.Edges {
		e = append(e, k)
	}
	sort.Sort(e)

	unionFindSet := CreateUnionFindSet(graph)
	var mstRes []*Edge
	for index := range e {
		from := e[index].From
		to := e[index].To
		if !unionFindSet.IsSameSet(from, to) {
			unionFindSet.Union(from, to)
			mstRes = append(mstRes, e[index])
		}
	}
	return mstRes
}
