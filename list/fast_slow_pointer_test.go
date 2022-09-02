package list

import (
	"reflect"
	"testing"
)

func TestFindNodeByFastSlowPointer1(t *testing.T) {
	type args struct {
		head1 *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fastSlowPointer1 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
			want: 1,
		},
		{
			name: "fastSlowPointer1 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next:  nil,
					},
				},
			},
			want: 1,
		},
		{
			name: "fastSlowPointer1 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "fastSlowPointer1 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 3,
							Next: &SingleLinkedNode{
								Value: 4,
								Next:  nil,
							},
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "fastSlowPointer1 test",
			args: args{
				head1: &SingleLinkedNode{
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
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNodeByFastSlowPointer1(tt.args.head1); !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("FindNodeByFastSlowPointer1() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestFindNodeByFastSlowPointer2(t *testing.T) {
	type args struct {
		head1 *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fastSlowPointer2 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
			want: 1,
		},
		{
			name: "fastSlowPointer2 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next:  nil,
					},
				},
			},
			want: 2,
		},
		{
			name: "fastSlowPointer2 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "fastSlowPointer2 test",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 3,
							Next: &SingleLinkedNode{
								Value: 4,
								Next:  nil,
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "fastSlowPointer2 test",
			args: args{
				head1: &SingleLinkedNode{
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
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNodeByFastSlowPointer2(tt.args.head1); !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("FindNodeByFastSlowPointer2() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}
