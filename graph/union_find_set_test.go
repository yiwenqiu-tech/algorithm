package graph

import (
	"testing"
)

func TestUnionFindSet_Union(t *testing.T) {
	a := &Node{}
	b := &Node{}
	c := &Node{}
	graph := Graph{
		Nodes: map[int]*Node{
			1: a,
			2: b,
			3: c,
		},
	}
	set := CreateUnionFindSet(&graph)
	res := set.IsSameSet(a, b)
	if res != false {
		t.Errorf("IsSameSet() = %v, want %v", res, false)
	}
	set.Union(a, b)
	res = set.IsSameSet(a, b)
	if res != true {
		t.Errorf("IsSameSet() = %v, want %v", res, true)
	}
}
