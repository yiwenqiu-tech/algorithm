package binarytree

import (
	"reflect"
	"testing"
)

func TestFindLowestAncestor1(t *testing.T) {
	type args struct {
		head  *Node
		node1 *Node
		node2 *Node
	}
	node1 := &Node{
		Value: 1,
	}
	node2 := &Node{
		Value: 2,
	}
	node3 := &Node{
		Value: 3,
	}
	node4 := &Node{
		Value: 3,
	}
	node5 := &Node{
		Value: 5,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "TestFindLowestAncestor1",
			args: args{
				head:  nil,
				node1: nil,
				node2: nil,
			},
			want: nil,
		},
		{
			name: "TestFindLowestAncestor2",
			args: args{
				head:  node1,
				node1: node4,
				node2: node4,
			},
			want: node4,
		},
		{
			name: "TestFindLowestAncestor3",
			args: args{
				head:  node1,
				node1: node2,
				node2: node4,
			},
			want: node2,
		},
		{
			name: "TestFindLowestAncestor4",
			args: args{
				head:  node1,
				node1: node4,
				node2: node5,
			},
			want: node2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLowestAncestor1(tt.args.head, tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindLowestAncestor1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLowestAncestor2(t *testing.T) {
	type args struct {
		head  *Node
		node1 *Node
		node2 *Node
	}
	node1 := &Node{
		Value: 1,
	}
	node2 := &Node{
		Value: 2,
	}
	node3 := &Node{
		Value: 3,
	}
	node4 := &Node{
		Value: 4,
	}
	node5 := &Node{
		Value: 5,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "TestFindLowestAncestor1",
			args: args{
				head:  nil,
				node1: nil,
				node2: nil,
			},
			want: nil,
		},
		{
			name: "TestFindLowestAncestor2",
			args: args{
				head:  node1,
				node1: node4,
				node2: node4,
			},
			want: node4,
		},
		{
			name: "TestFindLowestAncestor3",
			args: args{
				head:  node1,
				node1: node2,
				node2: node4,
			},
			want: node2,
		},
		{
			name: "TestFindLowestAncestor4",
			args: args{
				head:  node1,
				node1: node4,
				node2: node5,
			},
			want: node2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLowestAncestor2(tt.args.head, tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindLowestAncestor2() = %v, want %v", got, tt.want)
			}
		})
	}
}
