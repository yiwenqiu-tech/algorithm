package binarytree

import (
	"reflect"
	"testing"
)

func TestFindSuccessorNode1(t *testing.T) {
	type args struct {
		head *NodeWithParent
		node *NodeWithParent
	}
	node1 := &NodeWithParent{
		Value: 1,
	}
	node2 := &NodeWithParent{
		Value: 2,
	}
	node3 := &NodeWithParent{
		Value: 3,
	}
	node4 := &NodeWithParent{
		Value: 4,
	}
	node5 := &NodeWithParent{
		Value: 5,
	}
	node6 := &NodeWithParent{
		Value: 6,
	}
	node7 := &NodeWithParent{
		Value: 7,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Parent = node1
	node3.Parent = node1
	node2.Left = node4
	node2.Right = node5
	node4.Parent = node2
	node5.Parent = node2
	node3.Left = node6
	node3.Right = node7
	node6.Parent = node3
	node7.Parent = node3

	tests := []struct {
		name string
		args args
		want *NodeWithParent
	}{
		{
			name: "FindSuccessorNode1",
			args: args{
				head: node1,
				node: node5,
			},
			want: node1,
		},
		{
			name: "FindSuccessorNode2",
			args: args{
				head: node1,
				node: node6,
			},
			want: node3,
		},
		{
			name: "FindSuccessorNode3",
			args: args{
				head: node1,
				node: node7,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSuccessorNode1(tt.args.head, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSuccessorNode1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSuccessorNode2(t *testing.T) {
	type args struct {
		head *NodeWithParent
		node *NodeWithParent
	}
	node1 := &NodeWithParent{
		Value: 1,
	}
	node2 := &NodeWithParent{
		Value: 2,
	}
	node3 := &NodeWithParent{
		Value: 3,
	}
	node4 := &NodeWithParent{
		Value: 4,
	}
	node5 := &NodeWithParent{
		Value: 5,
	}
	node6 := &NodeWithParent{
		Value: 6,
	}
	node7 := &NodeWithParent{
		Value: 7,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Parent = node1
	node3.Parent = node1
	node2.Left = node4
	node2.Right = node5
	node4.Parent = node2
	node5.Parent = node2
	node3.Left = node6
	node3.Right = node7
	node6.Parent = node3
	node7.Parent = node3
	tests := []struct {
		name string
		args args
		want *NodeWithParent
	}{
		{
			name: "FindSuccessorNode1",
			args: args{
				head: node1,
				node: node5,
			},
			want: node1,
		},
		{
			name: "FindSuccessorNode2",
			args: args{
				head: node1,
				node: node6,
			},
			want: node3,
		},
		{
			name: "FindSuccessorNode3",
			args: args{
				head: node1,
				node: node7,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSuccessorNode2(tt.args.head, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSuccessorNode2() = %v, want %v", got, tt.want)
			}
		})
	}
}
