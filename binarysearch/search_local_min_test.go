package binarysearch

import "testing"

func TestSearchLocalMin(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only have a num",
			args: args{
				array: []int{1},
			},
			want: -1,
		},
		{
			name: "第一个数值比第二个小，则第一个已经是局部最小值",
			args: args{
				array: []int{1, 2},
			},
			want: 0,
		},
		{
			name: "最后一个数值比倒数二个小，则最后一个已经是局部最小值",
			args: args{
				array: []int{3, 2, 1},
			},
			want: 2,
		},
		{
			name: "局部最小值位于中间",
			args: args{
				array: []int{7, 5, 4, 3, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchLocalMin(tt.args.array); got != tt.want {
				t.Errorf("SearchLocalMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
