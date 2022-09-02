package list

import (
	"testing"
)

func TestPartitionByArray(t *testing.T) {
	type args struct {
		head *SingleLinkedNode
		num  int
	}
	tests := []struct {
		name string
		args args
		want *SingleLinkedNode
	}{
		{
			name: "test partition by array",
			args: args{
				head: &SingleLinkedNode{
					Value: 9,
					Next: &SingleLinkedNode{
						Value: 8,
						Next: &SingleLinkedNode{
							Value: 3,
							Next: &SingleLinkedNode{
								Value: 7,
								Next: &SingleLinkedNode{
									Value: 6,
									Next: &SingleLinkedNode{
										Value: 5,
										Next:  nil,
									},
								},
							},
						},
					},
				},
				num: 7,
			},
			want: &SingleLinkedNode{
				Value: 5,
				Next: &SingleLinkedNode{
					Value: 6,
					Next: &SingleLinkedNode{
						Value: 3,
						Next: &SingleLinkedNode{
							Value: 7,
							Next: &SingleLinkedNode{
								Value: 8,
								Next: &SingleLinkedNode{
									Value: 9,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartitionByArray(tt.args.head, tt.args.num)
			var nodeNum = 1
			for got != nil {
				if got.Value != tt.want.Value {
					t.Errorf("PartitionByArray() = %v, want %v, nodeNum: %v", got, tt.want, nodeNum)
				}
				got = got.Next
				tt.want = tt.want.Next
				nodeNum++
			}
		})
	}
}

func TestPartitionByPointer(t *testing.T) {
	type args struct {
		head *SingleLinkedNode
		num  int
	}
	tests := []struct {
		name string
		args args
		want *SingleLinkedNode
	}{
		{
			name: "test partition by array",
			args: args{
				head: &SingleLinkedNode{
					Value: 9,
					Next: &SingleLinkedNode{
						Value: 8,
						Next: &SingleLinkedNode{
							Value: 3,
							Next: &SingleLinkedNode{
								Value: 7,
								Next: &SingleLinkedNode{
									Value: 6,
									Next: &SingleLinkedNode{
										Value: 5,
										Next:  nil,
									},
								},
							},
						},
					},
				},
				num: 7,
			},
			want: &SingleLinkedNode{
				Value: 3,
				Next: &SingleLinkedNode{
					Value: 6,
					Next: &SingleLinkedNode{
						Value: 5,
						Next: &SingleLinkedNode{
							Value: 7,
							Next: &SingleLinkedNode{
								Value: 9,
								Next: &SingleLinkedNode{
									Value: 8,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test partition by array",
			args: args{
				head: &SingleLinkedNode{
					Value: 3,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 5,
							Next: &SingleLinkedNode{
								Value: 1,
								Next: &SingleLinkedNode{
									Value: 6,
									Next: &SingleLinkedNode{
										Value: 3,
										Next:  nil,
									},
								},
							},
						},
					},
				},
				num: 7,
			},
			want: &SingleLinkedNode{
				Value: 3,
				Next: &SingleLinkedNode{
					Value: 2,
					Next: &SingleLinkedNode{
						Value: 5,
						Next: &SingleLinkedNode{
							Value: 1,
							Next: &SingleLinkedNode{
								Value: 6,
								Next: &SingleLinkedNode{
									Value: 3,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test partition by array",
			args: args{
				head: &SingleLinkedNode{
					Value: 3,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 5,
							Next: &SingleLinkedNode{
								Value: 1,
								Next: &SingleLinkedNode{
									Value: 7,
									Next: &SingleLinkedNode{
										Value: 3,
										Next:  nil,
									},
								},
							},
						},
					},
				},
				num: 7,
			},
			want: &SingleLinkedNode{
				Value: 3,
				Next: &SingleLinkedNode{
					Value: 2,
					Next: &SingleLinkedNode{
						Value: 5,
						Next: &SingleLinkedNode{
							Value: 1,
							Next: &SingleLinkedNode{
								Value: 3,
								Next: &SingleLinkedNode{
									Value: 7,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test partition by array",
			args: args{
				head: &SingleLinkedNode{
					Value: 9,
					Next: &SingleLinkedNode{
						Value: 10,
						Next: &SingleLinkedNode{
							Value: 19,
							Next: &SingleLinkedNode{
								Value: 11,
								Next: &SingleLinkedNode{
									Value: 7,
									Next: &SingleLinkedNode{
										Value: 21,
										Next:  nil,
									},
								},
							},
						},
					},
				},
				num: 7,
			},
			want: &SingleLinkedNode{
				Value: 7,
				Next: &SingleLinkedNode{
					Value: 9,
					Next: &SingleLinkedNode{
						Value: 10,
						Next: &SingleLinkedNode{
							Value: 19,
							Next: &SingleLinkedNode{
								Value: 11,
								Next: &SingleLinkedNode{
									Value: 21,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartitionByPointer(tt.args.head, tt.args.num)
			var nodeNum = 1
			for got != nil {
				if got.Value != tt.want.Value {
					t.Errorf("PartitionByArray() = %v, want %v, nodeNum: %v", got, tt.want, nodeNum)
				}
				got = got.Next
				tt.want = tt.want.Next
				nodeNum++
			}
		})
	}
}
