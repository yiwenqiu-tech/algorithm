package graph

// UnionFindSet 并查集结构简化版
type UnionFindSet struct {
	Set map[*Node]*[]*Node
}

func CreateUnionFindSet(graph *Graph) *UnionFindSet {
	unionFindSet := &UnionFindSet{
		Set: make(map[*Node]*[]*Node),
	}
	for _, n := range graph.Nodes {
		unionFindSet.Set[n] = &[]*Node{n}
	}
	return unionFindSet
}

func (u *UnionFindSet) Union(n1 *Node, n2 *Node) {
	s1 := u.Set[n1]
	s2 := u.Set[n2]
	for index := range *s2 {
		*s1 = append(*s1, (*s2)[index])
		u.Set[(*s2)[index]] = s1
	}
}

func (u *UnionFindSet) IsSameSet(n1 *Node, n2 *Node) bool {
	if u.Set[n1] == u.Set[n2] {
		return true
	}
	return false
}
