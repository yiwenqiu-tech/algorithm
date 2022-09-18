package binarytree

import (
	"reflect"
	"testing"
)

func Test_bfs(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "BFS",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 4,
						},
						Right: &Node{
							Value: 5,
						},
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 6,
						},
						Right: &Node{
							Value: 7,
						},
					},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxWidth(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FindMaxWidthUseMap",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 4,
						},
						Right: &Node{
							Value: 5,
						},
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 6,
						},
						Right: &Node{
							Value: 7,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxWidthUseMap(tt.args.head); got != tt.want {
				t.Errorf("FindMaxWidthUseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxWidthNoUseMap(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FindMaxWidthNoUseMap",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 4,
						},
						Right: &Node{
							Value: 5,
						},
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 6,
						},
						Right: &Node{
							Value: 7,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxWidthNoUseMap(tt.args.head); got != tt.want {
				t.Errorf("FindMaxWidthNoUseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
