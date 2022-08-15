package list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		node *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want *SingleLinkedNode
	}{
		{
			name: "reverse list",
			args: args{
				&SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
			want: &SingleLinkedNode{
				Value: 1,
				Next:  nil,
			},
		},
		{
			name: "reverse list",
			args: args{
				&SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 3,
						Next:  nil,
					},
				},
			},
			want: &SingleLinkedNode{
				Value: 3,
				Next: &SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
		},
		{
			name: "reverse list",
			args: args{
				&SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 3,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 7,
								Next:  nil,
							},
						},
					},
				},
			},
			want: &SingleLinkedNode{
				Value: 7,
				Next: &SingleLinkedNode{
					Value: 10,
					Next: &SingleLinkedNode{
						Value: 3,
						Next: &SingleLinkedNode{
							Value: 1,
							Next:  nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseDoubleLinkedList(t *testing.T) {
	type args struct {
		node *DoubleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want *DoubleLinkedNode
	}{
		{
			name: "reverse double linked list",
			args: args{
				&DoubleLinkedNode{
					Value: 1,
					Next:  nil,
					Last:  nil,
				},
			},
			want: &DoubleLinkedNode{
				Value: 1,
				Next:  nil,
				Last:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseDoubleLinkedList(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseDoubleLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
