package list

import (
	"reflect"
	"testing"
)

func TestFindFirstIntersectNode(t *testing.T) {
	// 构造两条不相交且无环的单链表
	noLoopHead1 := &SingleLinkedNode{
		Value: 1,
		Next: &SingleLinkedNode{
			Value: 2,
			Next: &SingleLinkedNode{
				Value: 3,
				Next: &SingleLinkedNode{
					Value: 4,
					Next: &SingleLinkedNode{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}
	noLoopHead2 := &SingleLinkedNode{
		Value: 9,
		Next: &SingleLinkedNode{
			Value: 2,
			Next: &SingleLinkedNode{
				Value: 3,
				Next: &SingleLinkedNode{
					Value: 4,
					Next: &SingleLinkedNode{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}

	// 构造两个相交的无环单链表
	node5 := &SingleLinkedNode{
		Value: 5,
	}
	node4 := &SingleLinkedNode{
		Value: 4,
		Next:  node5,
	}
	node3 := &SingleLinkedNode{
		Value: 3,
		Next:  node4,
	}
	node2 := &SingleLinkedNode{
		Value: 2,
		Next:  node3,
	}
	node1 := &SingleLinkedNode{
		Value: 1,
		Next:  node2,
	}

	// 构造两个不相交的有环单链表
	nodeForLoop6 := &SingleLinkedNode{
		Value: 6,
	}
	nodeForLoop5 := &SingleLinkedNode{
		Value: 5,
	}
	nodeForLoop4 := &SingleLinkedNode{
		Value: 4,
		Next:  nodeForLoop6,
	}
	nodeForLoop3 := &SingleLinkedNode{
		Value: 3,
		Next:  nodeForLoop5,
	}
	nodeForLoop2 := &SingleLinkedNode{
		Value: 2,
		Next:  nodeForLoop4,
	}
	nodeForLoop1 := &SingleLinkedNode{
		Value: 1,
		Next:  nodeForLoop3,
	}
	nodeForLoop5.Next = nodeForLoop3
	nodeForLoop6.Next = nodeForLoop4

	// 构造两个相交的有环单链表(环外相交)
	nodeForLoop16 := &SingleLinkedNode{
		Value: 6,
	}
	nodeForLoop15 := &SingleLinkedNode{
		Value: 5,
	}
	nodeForLoop14 := &SingleLinkedNode{
		Value: 4,
		Next:  nodeForLoop15,
	}
	nodeForLoop13 := &SingleLinkedNode{
		Value: 3,
		Next:  nodeForLoop14,
	}
	nodeForLoop12 := &SingleLinkedNode{
		Value: 2,
		Next:  nodeForLoop13,
	}
	nodeForLoop11 := &SingleLinkedNode{
		Value: 1,
		Next:  nodeForLoop12,
	}
	nodeForLoop15.Next = nodeForLoop13
	nodeForLoop16.Next = nodeForLoop12

	// 构造两个相交的有环单链表(环内相交)
	nodeForLoop26 := &SingleLinkedNode{
		Value: 6,
	}
	nodeForLoop25 := &SingleLinkedNode{
		Value: 5,
	}
	nodeForLoop24 := &SingleLinkedNode{
		Value: 4,
		Next:  nodeForLoop25,
	}
	nodeForLoop23 := &SingleLinkedNode{
		Value: 3,
		Next:  nodeForLoop24,
	}
	nodeForLoop22 := &SingleLinkedNode{
		Value: 2,
		Next:  nodeForLoop23,
	}
	nodeForLoop21 := &SingleLinkedNode{
		Value: 1,
		Next:  nodeForLoop22,
	}
	nodeForLoop25.Next = nodeForLoop23
	nodeForLoop26.Next = nodeForLoop24

	type args struct {
		head1 *SingleLinkedNode
		head2 *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want *SingleLinkedNode
	}{
		{
			name: "two noLoop and no intersect list",
			args: args{
				head1: node1,
				head2: node4,
			},
			want: node4,
		},
		{
			name: "two noLoop and intersect list",
			args: args{
				head1: noLoopHead1,
				head2: noLoopHead2,
			},
			want: nil,
		},
		{
			name: "two noLoop and intersect list",
			args: args{
				head1: noLoopHead2,
				head2: noLoopHead1,
			},
			want: nil,
		},
		{
			name: "two Loop and no intersect list",
			args: args{
				head1: nodeForLoop1,
				head2: nodeForLoop2,
			},
			want: nil,
		},
		{
			name: "one Loop and one noLoop",
			args: args{
				head1: nodeForLoop1,
				head2: noLoopHead2,
			},
			want: nil,
		},
		{
			name: "two Loop and intersect of outer list",
			args: args{
				head1: nodeForLoop11,
				head2: nodeForLoop16,
			},
			want: nodeForLoop12,
		},
		{
			name: "two Loop and intersect of outer list",
			args: args{
				head1: nodeForLoop11,
				head2: nodeForLoop16,
			},
			want: nodeForLoop12,
		},
		{
			name: "two Loop and intersect of outer list",
			args: args{
				head1: nodeForLoop21,
				head2: nodeForLoop26,
			},
			want: nodeForLoop23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstIntersectNode(tt.args.head1, tt.args.head2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFirstIntersectNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
