package binarytree

import "testing"

func TestCheckSBT(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test check sbt0",
			args: args{
				head: nil,
			},
			want: true,
		},
		{
			name: "test check sbt1",
			args: args{
				head: &Node{
					Value: 1,
				},
			},
			want: true,
		},
		{
			name: "test check sbt2",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "test check sbt3",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "test check sbt4",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "test check sbt5",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "test check sbt6",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
					Left: &Node{
						Value: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "test check sbt7",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
					Left: &Node{
						Value: 3,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSBT(tt.args.head); got != tt.want {
				t.Errorf("CheckSBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckBST2(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test check sbt0",
			args: args{
				head: nil,
			},
			want: true,
		},
		{
			name: "test check sbt1",
			args: args{
				head: &Node{
					Value: 1,
				},
			},
			want: true,
		},
		{
			name: "test check sbt2",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "test check sbt3",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "test check sbt4",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "test check sbt5",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "test check sbt6",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
					Left: &Node{
						Value: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "test check sbt7",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
					Left: &Node{
						Value: 3,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSBT2(tt.args.head); got != tt.want {
				t.Errorf("CheckSBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCBT(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check cbt 1",
			args: args{
				head: nil,
			},
			want: true,
		},
		{
			name: "check cbt 2",
			args: args{
				head: &Node{
					Value: 1,
				},
			},
			want: true,
		},
		{
			name: "check cbt 3",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
				},
			},
			want: false,
		},
		{
			name: "check cbt 4",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
					Left: &Node{
						Value: 1,
					},
				},
			},
			want: true,
		},
		{
			name: "check cbt 5",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
					},
					Left: &Node{
						Value: 1,
						Left: &Node{
							Value: 1,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check cbt 6",
			args: args{
				head: &Node{
					Value: 1,
					Right: &Node{
						Value: 2,
						Left: &Node{
							Value: 1,
						},
					},
					Left: &Node{
						Value: 1,
						Left: &Node{
							Value: 1,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCBT(tt.args.head); got != tt.want {
				t.Errorf("CheckCBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckFBT(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check fbt 1",
			args: args{
				head: nil,
			},
			want: true,
		},
		{
			name: "check fbt 2",
			args: args{
				head: &Node{
					Value: 1,
				},
			},
			want: true,
		},
		{
			name: "check fbt 3",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "check fbt 4",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
					},
					Right: &Node{
						Value: 2,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFBT(tt.args.head); got != tt.want {
				t.Errorf("CheckFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckBBT(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check bbt 1",
			args: args{
				head: nil,
			},
			want: true,
		},
		{
			name: "check bbt 2",
			args: args{
				head: &Node{
					Value: 1,
				},
			},
			want: true,
		},
		{
			name: "check bbt 3",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
					},
				},
			},
			want: true,
		},
		{
			name: "check bbt 3",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
						Left: &Node{
							Value: 1,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBBT(tt.args.head); got != tt.want {
				t.Errorf("CheckBBT() = %v, want %v", got, tt.want)
			}
		})
	}
}
