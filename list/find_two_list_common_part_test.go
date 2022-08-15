package list

import (
	"reflect"
	"testing"
)

func TestFindTwoListCommonPart(t *testing.T) {
	type args struct {
		head1 *SingleLinkedNode
		head2 *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test find two list common part",
			args: args{
				head1: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 4,
						Next: &SingleLinkedNode{
							Value: 7,
							Next: &SingleLinkedNode{
								Value: 9,
								Next: &SingleLinkedNode{
									Value: 10,
									Next:  nil,
								},
							},
						},
					},
				},
				head2: &SingleLinkedNode{
					Value: 2,
					Next: &SingleLinkedNode{
						Value: 5,
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
			want: []int{7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTwoListCommonPart(tt.args.head1, tt.args.head2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTwoListCommonPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
