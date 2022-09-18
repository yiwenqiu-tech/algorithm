package binarytree

import (
	"reflect"
	"testing"
)

func TestRecursionOrder(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "recursion order",
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
			want: []int{1, 2, 4, 4, 4, 2, 5, 5, 5, 2, 1, 3, 6, 6, 6, 3, 7, 7, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			RecursionOrder(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("RecursionOrder result not match= %v, want %v", res, tt.want)
			}
		})
	}
}

func TestPreOrder(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pre order",
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
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			PreOrder(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("PreOrder result not match= %v, want %v", res, tt.want)
			}
		})
	}
}

func TestMidOrder(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "mid order",
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
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			MidOrder(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("MidOrder result not match= %v, want %v", res, tt.want)
			}
		})
	}
}

func TestPosOrder(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pos order",
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
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			PosOrder(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("PosOrder result not match= %v, want %v", res, tt.want)
			}
		})
	}
}

func TestPreOrderUnRecursive(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pre order unRecursive",
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
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "pre order unRecursive",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 4,
							Left: &Node{
								Value: 6,
							},
						},
						Right: &Node{
							Value: 5,
							Right: &Node{
								Value: 7,
							},
						},
					},
				},
			},
			want: []int{1, 2, 3, 4, 6, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			PreOrderUnRecursive(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("PreOrder unRecursive result not match= %v, want %v", res, tt.want)
			}
		})
	}
}

func TestPosOrderUnRecursive(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pos order unRecursive",
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
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "pos order unRecursive",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 4,
							Left: &Node{
								Value: 6,
							},
						},
						Right: &Node{
							Value: 5,
							Right: &Node{
								Value: 7,
							},
						},
					},
				},
			},
			want: []int{2, 6, 4, 7, 5, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			PosOrderUnRecursive(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("PosOrder unRecursive result not match= %v, want %v", res, tt.want)
			}
		})
	}
}

func TestMidOrderUnRecursive(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "mid order unRecursive",
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
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "pos order unRecursive",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 4,
							Left: &Node{
								Value: 6,
							},
						},
						Right: &Node{
							Value: 5,
							Right: &Node{
								Value: 7,
							},
						},
					},
				},
			},
			want: []int{2, 1, 6, 4, 3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			MidOrderUnRecursive(tt.args.head, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("MidOrder unRecursive result not match= %v, want %v", res, tt.want)
			}
		})
	}
}
